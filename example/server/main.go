package main

import (
	"log"
	"tcpx"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	srv := tcpx.NewServer()

	//定义tcp路由.
	srv.AddRouter(1000, Register)
	srv.AddRouter(1001, Login)

	if err := srv.Start("127.0.0.1:9999"); err != nil {
		panic(err)
	}
}

func Login(buf []byte) {
	log.Println("登录", string(buf))
}

func Register(buf []byte) {
	log.Println("注册", string(buf))
}
