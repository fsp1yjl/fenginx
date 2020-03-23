package fnet

import face "fenginx/finterface"

type Request struct {
	msg  face.IMessage
	conn face.IConnection
}

func (r *Request) Data() []byte {
	return r.msg.GetData()
}

func (r *Request) Connection() face.IConnection {
	return r.conn
}
