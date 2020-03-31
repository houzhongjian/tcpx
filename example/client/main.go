package main

import (
	"encoding/binary"
	"log"
	"net"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	conn, err := net.Dial("tcp4","127.0.0.1:9999")
	if err != nil {
		panic(err)
	}


	buf := Packet(1000, `{"account":"zhangsan", "password":"123456"}`)
	_, err = conn.Write(buf)
	if err != nil {
		log.Printf("err:%+v\n", err)
		return
	}

	buf = Packet(1001,`{"account":"zhangsan", "password":"123456"}`)
	_, err = conn.Write(buf)
	if err != nil {
		log.Printf("err:%+v\n", err)
		return
	}

	select {

	}
}

func Packet(id int, msg string) []byte {
	buf := make([]byte, 4+4+len(msg))

	binary.BigEndian.PutUint32(buf[0:4], uint32(id))
	binary.BigEndian.PutUint32(buf[4:8], uint32(len(msg)))
	copy(buf[8:], msg)

	return buf
}