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
	Routers   face.IRouters
}

func (c *Connection) Start() {
	defer fmt.Printf("connection handle func finishe, connid:,%d \n", c.ConnId)
	for {

		// buf := make([]byte, 512)

		conn := c.GetTCPConnection()
		// _, err := conn.Read(buf)
		// if err != nil {
		// 	if err == io.EOF {
		// 		break
		// 	}
		// 	fmt.Println("read error:", err)
		// 	continue
		// }

		p := &MsgPack{}
		// msg := &Message{}
		headerLen := p.HeaderLen()

		headerData := make([]byte, headerLen)

		if _, err := io.ReadFull(conn, headerData); err != nil {
			if err == io.EOF {
				fmt.Println("客户端断开连接")
				break
			}
			fmt.Println("read msg head error ", err)
			continue
		}

		fmt.Println("start header unpack....")
		msg, err := p.UnPack(headerData)
		if err != nil {
			fmt.Println("unpack msg header error: ", err)
			continue
		}

		fmt.Println("after header unpack....", msg.GetLength())

		msgData := make([]byte, msg.GetLength())
		if _, err := io.ReadFull(conn, msgData); err != nil {
			fmt.Println("read msg body error ", err)
			continue
		}
		msg.SetData(msgData)
		req := Request{
			msg:  msg,
			conn: c,
		}

		// go c.Router.handle(req)
		go func(req face.IRequest) {
			router := c.Routers.GetRouter(msg.GetID())
			if router != nil {
				router.Handle(req)
			}

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

func NewConnection(conn *net.TCPConn, connId int, r face.IRouters) face.IConnection {

	c := &Connection{
		ConnId:  connId,
		Conn:    conn,
		Routers: r,
	}

	return c

}
