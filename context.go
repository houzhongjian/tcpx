package tcpx

import (
	"encoding/binary"
	"io"
	"log"
	"net"
)

type Context struct {
	conn net.Conn
	Body []byte
}

//Close 断开连接.
func (c *Context) Close() {
	defer c.conn.Close()
}

//NewConn.
func (srv *Handler) NewConn(conn net.Conn) *Context {
	c := new(Context)
	c.conn = conn
	return c
}

//GetPacket 获取发送的数据包.
func (c *Context) GetPacket(router map[int]tcpxFunction) {
	log.Println("开始处理conn")

	for {
		//获取协议号.
		buf := make([]byte, 4)
		_, err := io.ReadFull(c.conn, buf)
		if err != nil {
			log.Printf("err:%+v\n", err)
			c.Close()
			return
		}
		protocol := int32(binary.BigEndian.Uint32(buf))

		//获取数据大小.
		buf = make([]byte, 4)
		_, err = io.ReadFull(c.conn, buf)
		if err != nil {
			log.Printf("err:%+v\n", err)
			c.Close()
			return
		}
		size := int32(binary.BigEndian.Uint32(buf))

		//获取数据内容.
		buf = make([]byte, size)
		_, err = io.ReadFull(c.conn, buf)
		if err != nil {
			log.Printf("err:%+v\n", err)
			c.Close()
			return
		}

		c.Body = buf

		//转发数据给对应的方法.
		funcName := router[int(protocol)]
		funcName(c)
	}
}

func (c *Context) Write(buf []byte) error {
	_, err := c.conn.Write(buf)
	return err
}
