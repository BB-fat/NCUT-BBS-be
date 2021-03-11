package controller

import (
	"github.com/gin-gonic/gin"
	"ncutbbs/module/session"
	"net/http"
)

func Test1(c *gin.Context) {

}

func Test2(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func Test3(c *gin.Context) {
	s := session.Get(c.GetHeader("token"))
	session.RefreshToken(s)
}
