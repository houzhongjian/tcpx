package tcpx

import (
	"encoding/binary"
	"io"
	"log"
)

//GetMessageID 获取消息id.
func (client *Client)GetMessageID() int {
	var buf = make([]byte, 4)
	_, err := io.ReadFull(client.conn, buf)
	if err != nil {
		log.Printf("err:%+v\n", err)
		client.Close()
	}

	binary.BigEndian.Uint32(buf)

	return 0
}

func (client *Client) Close() {
	defer client.conn.Close()
}