package gin_protobuf_json_converter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

var marshaler = jsonpb.Marshaler{}
var unmarshaler = jsonpb.Unmarshaler{}

func SendPBAsJSON(c *gin.Context, statusCode int, pbMsg proto.Message) {
	b, err := marshaler.MarshalToString(pbMsg)

	if err != nil {
		fmt.Println("error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "protobuf message couldn't be transform to json"})
	}
	c.Data(statusCode, "application/json; charset=utf-8", []byte(b))
}

func JsonToPb(c *gin.Context, pbObj interface{}) error {
	err := unmarshaler.Unmarshal(c.Request.Body, (pbObj).(proto.Message))
	return err
}

func SetMarshaler(m jsonpb.Marshaler) {
	marshaler = m
}

func SetUnMarshaler(um jsonpb.Unmarshaler) {
	unmarshaler = um
}
