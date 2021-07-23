package api

import (
	"fmt"
	"os"
	"time"

	"github.com/maldan/gam-app-backuper/internal/app/backuper/core"
	"github.com/maldan/go-cmhp/cmhp_crypto"
	"github.com/maldan/go-cmhp/cmhp_file"
	"github.com/maldan/go-cmhp/cmhp_process"
	"github.com/maldan/go-cmhp/cmhp_time"
	"github.com/maldan/go-restserver"
)

type MainApi struct {
}

func (r MainApi) GetIndex() string {
	return "Test"
}

// Load config
func (r MainApi) GetConfig() core.Config {
	var config core.Config
	cmhp_file.ReadJSON(core.DataDir+"/config.json", &config)
	return config
}

// Save config
func (r MainApi) PostConfig(args core.Config) {
	cmhp_file.WriteJSON(core.DataDir+"/config.json", &args)
}

// Get backup list
func (r MainApi) GetBackupList(args ArgsBackup) []string {
	return core.List("backup/gam-data/" + args.Path)
}

// Backup app data
func (r MainApi) PostBackup(args ArgsBackup) {
	// Zip
	zipFile := os.TempDir() + "/" + cmhp_crypto.UID(10) + ".zip"
	p, _ := cmhp_process.Create("zip", "-9", "-r", zipFile, ".", "-i", "*")
	p.Dir = core.DataDir + "/../" + args.Path
	p.Run()
	defer cmhp_file.Delete(zipFile)

	data, err := cmhp_file.ReadBin(zipFile)
	if err != nil {
		restserver.Fatal(500, restserver.ErrorType.Unknown, "path", err.Error())
	}

	// Upload
	core.Upload(
		"backup/gam-data/"+args.Path+"/"+cmhp_time.Format(time.Now(), "YYYY-MM-DD HH_mm_ss")+".zip",
		data,
		"public-read",
		"application/zip",
	)
}

// Get file list
func (r MainApi) GetList() []string {
	out := make([]string, 0)
	files, _ := cmhp_file.List(core.DataDir + "/../")
	fmt.Println(core.DataDir + "/../")
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		out = append(out, file.Name())
	}
	return out
}
