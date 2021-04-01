package main

import (
	"ncutbbs/config"
	"ncutbbs/model"
	"ncutbbs/module/oss"
	"ncutbbs/server"
)

func main() {
	config.ParseConfig()
	model.InitDB()
	oss.InitOSS()
	server.Run()
}
