package main

import (
	"embed"

	helloworld "github.com/maldan/gam-app-backuper/internal/app/backuper"
)

//go:embed frontend/build/*
var frontFs embed.FS

func main() {
	helloworld.Start(frontFs) //
}
