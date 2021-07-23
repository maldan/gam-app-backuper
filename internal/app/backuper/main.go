package backuper

import (
	"embed"
	"flag"
	"fmt"
	"os"

	"github.com/maldan/gam-app-backuper/internal/app/backuper/api"
	"github.com/maldan/gam-app-backuper/internal/app/backuper/core"
	"github.com/maldan/go-restserver"
	"github.com/zserge/lorca"
)

func Start(frontFs embed.FS) {
	var host = flag.String("host", "127.0.0.1", "Server Hostname")
	var port = flag.Int("port", 16000, "Server Port")
	_ = flag.Int("clientPort", 8080, "Client Port")
	var gui = flag.Bool("gui", false, "Use Gui")
	var width = flag.Int("width", 1100, "Window Width")
	var height = flag.Int("height", 900, "Window Height")
	var dataDir = flag.String("dataDir", "db", "Data Directory")
	_ = flag.String("appId", "id", "App id")
	flag.Parse()
	core.DataDir = *dataDir

	if *gui {
		go (func() {
			ui, _ := lorca.New("", "", *width, *height)
			defer ui.Close()
			ui.Load(fmt.Sprintf("http://%s:%d/", *host, *port))
			<-ui.Done()
			os.Exit(0)
		})()
	}

	core.StartS3()

	restserver.Start(fmt.Sprintf("%s:%d", *host, *port), map[string]interface{}{
		"/": restserver.VirtualFs{Root: "frontend/build/", Fs: frontFs},
		"/api": map[string]interface{}{
			"main": api.MainApi{},
		},
	})
}
