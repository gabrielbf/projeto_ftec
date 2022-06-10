package main

import (
	"os"

	"github.com/sample-project/internal/app"
)

// @title Sample Project
// @version 1.0
// @description This doc presents the http endpoints that could be used for manage samples.

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	rootdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	app.Init(rootdir)
	ch := make(chan bool, 1)
	<-ch
}
