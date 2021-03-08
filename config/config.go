package config

import "flag"

type ServerConfig struct {
	Host  string
	Port  int
	Debug bool
}

var Config ServerConfig

func ParseConfig() {
	Config = ServerConfig{}
	flag.StringVar(&Config.Host, "Host", "localhost", "Host")
	flag.IntVar(&Config.Port, "Port", 8080, "端口")
	flag.BoolVar(&Config.Debug, "Debug", true, "debug模式")
	flag.Parse()
}
