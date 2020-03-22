package face

type IServer interface {
	Start()

	Stop()

	Serve()

	AddRouter(r IRouter)
}
