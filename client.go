package tcpx

import (
	"encoding/binary"
	"io"
	"log"
	"net"
)

type Client struct {
	conn net.Conn
}

//NewClient.
func (srv *Handler) NewClient(conn net.Conn) *Client {
	client := new(Client)
	client.conn = conn
	return client
}

//GetPacket 获取发送的数据包.
func (client *Client) GetPacket(router map[int]function) {
	log.Println("开始处理conn")
	defer client.conn.Close()
	for {
		//获取协议号.
		buf := make([]byte, 4)
		_, err := io.ReadFull(client.conn, buf)
		if err != nil {
			log.Printf("err:%+v\n",err)
			return
		}
		protocol := int32(binary.BigEndian.Uint32(buf))

		//获取数据大小.
		buf = make([]byte, 4)
		_, err = io.ReadFull(client.conn, buf)
		if err != nil {
			log.Printf("err:%+v\n",err)
			return
		}
		size := int32(binary.BigEndian.Uint32(buf))

		//获取数据内容.
		buf = make([]byte, size)
		_, err = io.ReadFull(client.conn, buf)
		if err != nil {
			log.Printf("err:%+v\n",err)
			return
		}

		//转发数据给对应的方法.
		funcName := router[int(protocol)]
		funcName(buf)
	}
}