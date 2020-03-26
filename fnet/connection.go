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
	MsgChan   chan []byte
}

func (c *Connection) StartReader() {
	for {

		conn := c.GetTCPConnection()

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

func (c *Connection) StartWrite() {

	defer fmt.Println(c.RemoteAddr().String(), " conn Writer exit!")
	for {
		select {
		case data := <-c.MsgChan:
			//有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send Data error:, ", err, " Conn Writer exit")
				return
			}
		}
	}
}

func (c *Connection) Start() {
	defer fmt.Printf("connection handle func finishe, connid:,%d \n", c.ConnId)
	go c.StartReader()
	go c.StartWrite()

}

func (c *Connection) Stop() {

}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnId() int {
	return c.ConnId
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) SendMsg(msgID uint32, data []byte) {

	//先对数据进行封包
	msg := NewMessage(msgID, data)
	p := &MsgPack{}
	msgByte, _ := p.Pack(msg)

	c.MsgChan <- msgByte
}

func NewConnection(conn *net.TCPConn, connId int, r face.IRouters) face.IConnection {

	c := &Connection{
		ConnId:  connId,
		Conn:    conn,
		Routers: r,
		MsgChan: make(chan []byte),
	}

	return c

}
