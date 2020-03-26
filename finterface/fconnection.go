package face

import (
	"net"
)

type IConnection interface {
	GetTCPConnection() *net.TCPConn
	GetConnId() int
	RemoteAddr() net.Addr
	SendMsg(msgID uint32, data []byte)
	Start()
	Stop()
}

type HandleFunc func(*net.TCPConn, []byte, int) error
