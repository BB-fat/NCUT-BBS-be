package config

import "flag"

type ServerConfig struct {
	Host  string
	Port  int
	Debug bool
	//数据库
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
	// session有效时间
	SessionTTL int
	// oss参数
	OSSAccessKeyID     string
	OSSAccessKeySecret string
	OSSEndpoint        string
	OSSBaseURL         string
}

var Config ServerConfig

func ParseConfig() {
	Config = ServerConfig{}
	//服务配置
	flag.StringVar(&Config.Host, "Host", "localhost", "Host")
	flag.IntVar(&Config.Port, "Port", 8080, "端口")
	flag.BoolVar(&Config.Debug, "Debug", true, "debug模式")
	//数据库配置
	flag.StringVar(&Config.DBUser, "DBUser", "", "数据库用户名")
	flag.StringVar(&Config.DBPassword, "DBPassword", "", "数据库密码")
	flag.StringVar(&Config.DBHost, "DBHost", "", "数据库host")
	flag.IntVar(&Config.DBPort, "DBPort", 3306, "数据库端口")
	flag.StringVar(&Config.DBName, "DBName", "ncut_bbs", "数据库名")
	// session 有效时间 默认两小时
	flag.IntVar(&Config.SessionTTL, "SessionTTL", 60*60*2, "session有效时间")

	// OSS
	flag.StringVar(&Config.OSSEndpoint, "OSSEndpoint", "", "OSSEndpoint")
	flag.StringVar(&Config.OSSBaseURL, "OSSBaseURL", "", "OSSBaseURL")
	flag.StringVar(&Config.OSSAccessKeyID, "OSSAccessKeyID", "", "OSSAccessKeyID")
	flag.StringVar(&Config.OSSAccessKeySecret, "OSSAccessKeySecret", "", "OSSAccessKeySecret")

	flag.Parse()
}
