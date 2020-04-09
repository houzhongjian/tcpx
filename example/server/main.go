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
	srv.AddRouter(1002, Index)
	srv.AddRouter(1003, LogOut)

	if err := srv.Start("127.0.0.1:9999"); err != nil {
		panic(err)
	}
}

func Login(c *tcpx.Context) {
	log.Println("登录", string(c.Body))

	type Account struct {
		Account string	`json:"account"`
		Password string	`json:"password"`
	}

	msg := Account{}
	if err := c.BindJSON(&msg); err != nil {
		log.Printf("err:%+v\n",err)
		return
	}
	log.Printf("msg:%+v\n", msg)
	tcpx.GetPlayerManager().Set(123, "account",msg.Account)
	tcpx.GetPlayerManager().Set(123, "password",msg.Password)
}

func Index(c *tcpx.Context) {
	log.Println("index")
	log.Println(tcpx.GetPlayerManager().Get(123))
}


func LogOut(c *tcpx.Context) {
	log.Println("退出登录", string(c.Body))
	log.Println(tcpx.GetPlayerManager().Get(123,"face").(string))
}

func Register(c *tcpx.Context) {
	log.Println("注册", string(c.Body))
}
