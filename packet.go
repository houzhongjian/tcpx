package tcpx

import (
	"encoding/binary"
	"io"
	"log"
)

//GetMessageID 获取消息id.
func (c *Context) GetMessageID() int {
	var buf = make([]byte, 4)
	_, err := io.ReadFull(c.conn, buf)
	if err != nil {
		log.Printf("err:%+v\n", err)
		c.Close()
	}

	binary.BigEndian.Uint32(buf)

	return 0
}
