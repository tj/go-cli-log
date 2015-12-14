package main

import "github.com/tj/go-cli-log"
import "errors"
import "fmt"

func main() {
	fmt.Println()
	log.Verbose = true
	log.Info("install", "package %s@%s", "express", "3.2.1")
	log.Debug("fetch", "tarball for express")
	log.Info("unpack", "tarball to node_modules/express")
	log.Debug("unpack", "express/Readme.md")
	log.Debug("unpack", "express/lib/application.js")
	log.Debug("unpack", "express/lib/request.js")
	log.Debug("unpack", "express/lib/response.js")
	log.Warn("duplicate %s package found", "express@3.x")
	log.Error(errors.New("something exploded"))
	fmt.Println()
}
