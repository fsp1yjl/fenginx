package main

import (
	"fenginx/fnet"
	"fmt"
	"net"
)

func main() {

	conn, _ := net.Dial("tcp4", "0.0.0.0:8889")

	msg := fnet.Message{}
	msg.SetID(2)
	msg.SetData([]byte("hello,eric"))
	len := uint32(len(msg.GetData()))
	fmt.Println("lenlllll--:", len)
	msg.SetLength(len)

	p := fnet.MsgPack{}
	sendBuf, err := p.Pack(&msg)
	if err != nil {
		fmt.Println(" client message pack error", err)
		panic("hhh")
	}

	conn.Write(sendBuf)

	buf := make([]byte, 512)
	cnt, _ := conn.Read(buf)
	fmt.Printf(" server call back : %s, cnt = %d\n", buf, cnt)

}
