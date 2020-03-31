package tcpx

import (
	"log"
	"net"
)

type TcpServer interface {
	AddRouter(msgid int, handle func([]byte))
	Start(addr string) error
}

type function func(buf []byte)


type Handler struct {
	addr string
	clientList chan *Client
	router map[int]function
}

func NewServer() TcpServer {
	srv := new(Handler)
	srv.clientList = make(chan *Client)
	srv.router = make(map[int]function)
	return srv
}

//handler 处理具体的连接.
func (srv *Handler) handler() {
	for {
		select {
		case client := <-srv.clientList:
			go client.GetPacket(srv.router)
		}
	}
}

func (srv *Handler) AddRouter(msgid int, handle func(buf []byte)) {
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
			log.Printf("err:%+v\n",err)
			continue
		}

		log.Println(conn)
		//获取当前的协议号，转发给不同的方法实现.
		client := srv.NewClient(conn)
		srv.clientList <- client
	}
}