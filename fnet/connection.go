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
	Router    face.IRouter
}

func (c *Connection) Start() {
	defer fmt.Printf("connection handle func finishe, connid:,%d \n", c.ConnId)
	for {

		buf := make([]byte, 512)

		conn := c.GetTCPConnection()
		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			continue
		}

		req := Request{
			data: buf,
			conn: c,
		}

		// go c.Router.handle(req)
		go func(req face.IRequest) {
			c.Router.Handle(req)
		}(&req)
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

func NewConnection(conn *net.TCPConn, connId int, r face.IRouter) face.IConnection {

	c := &Connection{
		ConnId: connId,
		Conn:   conn,
		Router: r,
	}

	return c

}
