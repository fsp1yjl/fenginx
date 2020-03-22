package fnet

import (
	face "fenginx/finterface"
	"fmt"
	"io"
	"net"
)

type Connection struct {
	ConnId    int
	Conn      *net.TCPConn
	HandleAPI face.HandleFunc
}

func (c *Connection) Start() {
	defer fmt.Printf("connection handle func finishe, connid:,%d \n", c.ConnId)
	for {

		buf := make([]byte, 512)

		conn := c.GetTCPConnection()
		cnt, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			continue
		}

		c.HandleAPI(c.Conn, buf, cnt)
	}

}

func (c *Connection) Stop() {

}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnId() int {
	return c.ConnId
}

func (c *Connection) RemoteAddr() {
}

func NewConnection(conn *net.TCPConn, connId int, callbackAPI face.HandleFunc) face.IConnection {

	c := &Connection{
		ConnId:    connId,
		Conn:      conn,
		HandleAPI: callbackAPI,
	}

	return c

}
