package tcpx

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
)


//BindJSON 绑定json.
func (c *Context) BindJSON(obj interface{}) error {
	return json.Unmarshal(c.Body, &obj)
}

func (c *Context) BindPb(obj proto.Message) error {
	return proto.Unmarshal(c.Body, obj)
}

