package main

import (
	"ncutbbs/config"
	"ncutbbs/server"
)

func main() {
	config.ParseConfig()
	server.Run()
}
