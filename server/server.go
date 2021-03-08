package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ncutbbs/config"
)

func Run() {
	r := gin.Default()
	setUrl(r)
	_ = r.Run(fmt.Sprintf("%s:%d", config.Config.Host, config.Config.Port))
}
