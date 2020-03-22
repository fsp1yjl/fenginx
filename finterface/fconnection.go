package face

import (
	"net"
)

type IConnection interface {
	GetTCPConnection() *net.TCPConn
	GetConnId() int
	RemoteAddr()
	// Send()
	Start()
	Stop()
}

type HandleFunc func(*net.TCPConn, []byte, int) error
