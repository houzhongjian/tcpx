package tcpx

import (
	"testing"
)

func TestNew(t *testing.T) {
	srv := NewServer()
	srv.AddRouter(1000, Heartbeat) //心跳.
	srv.AddRouter(1001, Login)     //登录协议.
	srv.Start("127.0.0.1:9999")
}
