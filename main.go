package main

import (
	"embed"

	helloworld "github.com/maldan/gam-app-rapi_debug/internal/app/rapi_debug"
)

//go:embed frontend/build/*
var frontFs embed.FS

func main() {
	helloworld.Start(frontFs) //
}
