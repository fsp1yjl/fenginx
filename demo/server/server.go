package main

import (
	face "fenginx/finterface"
	fnet "fenginx/fnet"
	utils "fenginx/utils"
)

type TestRouter struct {
	fnet.BaseRouter
}

func (t *TestRouter) Handle(req face.IRequest) {

	c := req.Connection()

	d := req.Data()

	// 如下4步骤可以封装到connection struct中
	// msg := fnet.NewMessage(req.MsgID(), d)
	// p := &fnet.MsgPack{}
	// msgByte, _ := p.Pack(msg)
	// c.MsgChan <- msgByte
	c.SendMsg(req.MsgID(), d)
}

type Test2Router struct {
	fnet.BaseRouter
}

func (t *Test2Router) Handle(req face.IRequest) {

	c := req.Connection()
	// c := (req.Connection()).(*fnet.Connection)

	c.SendMsg(req.MsgID(), req.Data())

}

func main() {
	//加载配置文件
	utils.G.LoadConfig()

	s := fnet.NewServer()
	s.AddRouter(1, &TestRouter{})
	s.AddRouter(2, &Test2Router{})
	s.Serve()
}
