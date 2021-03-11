package tool

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func LoadMessage(c *gin.Context, m proto.Message) {
	body, _ := c.GetRawData()
	_ = protojson.Unmarshal(body, m)
}

func DumpMessage(m proto.Message) string {
	json, _ := protojson.Marshal(m)
	return string(json)
}
