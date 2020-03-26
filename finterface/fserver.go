package face

type IServer interface {
	Start()

	Stop()

	Serve()

	AddRouter(msgId uint32, r IRouter)
}
