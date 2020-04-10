package tcpx

import (
	"log"
	"net"
)

type TcpServer interface {
	AddRouter(msgid int, handle func(c *Context))
	Start(addr string) error
}

type tcpxFunction func(ctx *Context)

type Handler struct {
	addr       string
	clientConn chan *Context
	router     map[int]tcpxFunction
	maxConn    int
	session    int
}

func NewServer() TcpServer {
	srv := new(Handler)
	srv.clientConn = make(chan *Context)
	srv.router = make(map[int]tcpxFunction)
	return srv
}

//SetMaxConn 设置最大的tcp连接数.
func (srv *Handler) SetMaxConn(num int) {
	srv.maxConn = num
}

//handler 处理具体的连接.
func (srv *Handler) handler() {
	for {
		select {
		case clientConn := <-srv.clientConn:
			go clientConn.GetPacket(srv.router)
		}
	}
}

func (srv *Handler) AddRouter(msgid int, handle func(ctx *Context)) {
	if _, ok := srv.router[msgid]; ok {
		panic("当前消息id已经存在")
	}
	srv.router[msgid] = handle
}

func (srv *Handler) Start(addr string) error {
	srv.addr = addr

	go srv.handler()
	return srv.start()
}

func (srv *Handler) start() error {
	listener, err := net.Listen("tcp4", srv.addr)
	if err != nil {
		return err
	}

	srv.accept(listener)
	return nil
}

func (srv *Handler) accept(listener net.Listener) {
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("err:%+v\n", err)
			continue
		}
		if srv.session == srv.maxConn {
			log.Println("超过系统最大连接数")
			conn.Close()
			continue
		}

		//获取当前的协议号，转发给不同的方法实现.
		clientConn := srv.NewConn(conn)
		srv.clientConn <- clientConn
		srv.session += 1
	}
}
