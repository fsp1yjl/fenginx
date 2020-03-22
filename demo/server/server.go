package main

import (
	face "fenginx/finterface"
	fnet "fenginx/fnet"
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
	s := fnet.NewServer("test server")
	s.AddRouter(&TestRouter{})
	s.Serve()
}
