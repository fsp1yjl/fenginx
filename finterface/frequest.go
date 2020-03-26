package face

type IRequest interface {
	Connection() IConnection //获取request对应的连接信息
	Data() []byte            //获取请求中的数据
	MsgID() uint32
}
