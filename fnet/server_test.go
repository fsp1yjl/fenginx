package fnet

import (
	face "fenginx/finterface"
	utils "fenginx/utils"
	"fmt"
	"net"
	"testing"
	"time"
)

func ClientTest() {
	conn, err := net.Dial("tcp4", "localhost:8889")
	if err != nil {
		fmt.Println("dial erroor:", err)
		return
	}

	msg := Message{}
	msg.SetID(1)
	msg.SetData([]byte("hello,eric"))
	len := uint32(len(msg.GetData()))
	fmt.Println("lenlllll--:", len)
	msg.SetLength(len)

	p := MsgPack{}
	sendBuf, err := p.Pack(&msg)
	if err != nil {
		fmt.Println(" client message pack error", err)
		panic("hhh")
	}

	conn.Write(sendBuf)
	buf := make([]byte, 512)
	cnt, _ := conn.Read(buf)
	fmt.Printf(" server response : %s, cnt = %d\n", buf, cnt)

	time.Sleep(1 * time.Second)
	return
}

type TestRouter struct {
	BaseRouter
}

func (t *TestRouter) Handle(req face.IRequest) {

	c := req.Connection().GetTCPConnection()
	d := req.Data()
	if _, err := c.Write(d); err != nil {
		fmt.Println("write error:", err)
	}

}

func Test(t *testing.T) {

	/*
		服务端测试
	*/
	// 加载配置文件
	utils.G.LoadConfig()

	// 创建一个server 句柄 s
	s := NewServer()
	s.AddRouter(&TestRouter{})
	/*
		客户端测试
	*/
	go func() {
		fmt.Println("start client testing")
		time.Sleep(2 * time.Second)
		ClientTest()
	}()

	//2 开启服务
	s.Serve()
}
