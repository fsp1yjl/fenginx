package face

type IRouters interface {
	SetRouter(msgID uint32, r IRouter)
	GetRouter(msgID uint32) IRouter
	WorkerPoolStart()
	SendRequestToWorker(req IRequest)
	RequestHandle(req IRequest)
	GetWorkerPoolCount() uint32
}
