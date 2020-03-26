package main

import (
	face "fenginx/finterface"
	fnet "fenginx/fnet"
	utils "fenginx/utils"
	"fmt"
)

type TestRouter struct {
	fnet.BaseRouter
}

func (t *TestRouter) Handle(req face.IRequest) {

	c := req.Connection().GetTCPConnection()
	d := req.Data()
	if _, err := c.Write(d); err != nil {
		fmt.Println("write error:", err)
	}

}

type Test2Router struct {
	fnet.BaseRouter
}

func (t *Test2Router) Handle(req face.IRequest) {

	c := req.Connection().GetTCPConnection()
	d := req.Data()
	dd := []byte("test2 router.....")
	d = append(d, dd...)
	if _, err := c.Write(d); err != nil {
		fmt.Println("write error:", err)
	}

}

func main() {
	//加载配置文件
	utils.G.LoadConfig()

	s := fnet.NewServer()
	s.AddRouter(1, &TestRouter{})
	s.AddRouter(2, &Test2Router{})
	s.Serve()
}
