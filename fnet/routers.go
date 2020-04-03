package fnet

import (
	face "fenginx/finterface"
	"fmt"
)

type Routers struct {
	routers        map[uint32]face.IRouter
	workerCount    uint32
	workerReqQueue []chan face.IRequest
}

func NewRouters() face.IRouters {
	var workerCnt uint32 = 5 //todo:后期改为通过全局变量获取
	return &Routers{
		routers:        make(map[uint32]face.IRouter, 0),
		workerCount:    workerCnt,
		workerReqQueue: make([]chan face.IRequest, workerCnt),
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

// 将每次一个客户端请求封装为request对象，放入对应worker的队列中，等待worker做后续处理
func (rs *Routers) SendRequestToWorker(r face.IRequest) {
	var workerCnt uint32 = 5
	connID := uint32(r.Connection().GetConnId())
	workerID := connID % workerCnt
	rs.workerReqQueue[workerID] <- r
}

func (rs *Routers) RequestHandle(req face.IRequest) {
	msgID := req.MsgID()
	router := rs.GetRouter(msgID)
	router.Handle(req)
}

func (rs *Routers) OneWorkerRun(workerID uint32) {
	c := rs.workerReqQueue[workerID]

	for {
		req := <-c
		fmt.Println("worker index ::::", workerID)
		rs.RequestHandle(req)
	}
}

func (rs *Routers) WorkerPoolStart() {
	var workerCnt uint32 = 5
	var workerBacklog uint32 = 1000
	var index uint32
	for index = 0; index < workerCnt; index++ {
		// 启动时候指定每个worker任务队列的长度，后期考虑此步骤放在routers初始化的时候做
		rs.workerReqQueue[index] = make(chan face.IRequest, workerBacklog)
		go rs.OneWorkerRun(index)
	}
}

func (rs *Routers) GetWorkerPoolCount() uint32 {
	return rs.workerCount
}
