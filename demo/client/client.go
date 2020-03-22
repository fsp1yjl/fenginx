package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	conn, _ := net.Dial("tcp4", "0.0.0.0:8889")

	conn.Write([]byte("hello world"))

	buf := make([]byte, 512)
	cnt, _ := conn.Read(buf)

	fmt.Printf(" server call back : %s, cnt = %d\n", buf, cnt)

	time.Sleep(2 * time.Second)

	conn.Write([]byte("hello world, RETRY"))

	cnt, _ = conn.Read(buf)
	fmt.Printf(" server call back : %s, cnt = %d\n", buf, cnt)

}
