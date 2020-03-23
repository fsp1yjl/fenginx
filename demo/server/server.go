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

func main() {
	//加载配置文件
	utils.G.LoadConfig()

	s := fnet.NewServer()
	s.AddRouter(&TestRouter{})
	s.Serve()
}
