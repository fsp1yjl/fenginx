package fnet

import face "fenginx/finterface"

type Routers struct {
	routers map[uint32]face.IRouter
}

func NewRouters() face.IRouters {
	return &Routers{
		routers: make(map[uint32]face.IRouter, 0),
	}
}

func (rs *Routers) SetRouter(msgID uint32, r face.IRouter) {
	// todo 后期考虑rs.routers 为nil的情况
	rs.routers[msgID] = r
}

func (rs *Routers) GetRouter(msgID uint32) face.IRouter {
	// todo 后期考虑rs.routers 为nil的情况
	return rs.routers[msgID]
}
