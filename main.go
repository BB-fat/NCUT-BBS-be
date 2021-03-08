package main

import (
	"ncutbbs/config"
	"ncutbbs/model"
	"ncutbbs/server"
)

func main() {
	config.ParseConfig()
	model.InitDB()
	server.Run()
}
