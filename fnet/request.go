package fnet

import face "fenginx/finterface"

type Request struct {
	data []byte
	conn face.IConnection
}

func (r *Request) Data() []byte {
	return r.data
}

func (r *Request) Connection() face.IConnection {
	return r.conn
}
