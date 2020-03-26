package fnet

import face "fenginx/finterface"

type Request struct {
	msg  face.IMessage
	conn face.IConnection
}

func (r *Request) Data() []byte {
	return r.msg.GetData()
}

func (r *Request) MsgID() uint32 {
	return r.msg.GetID()
}

func (r *Request) Connection() face.IConnection {
	return r.conn
}
