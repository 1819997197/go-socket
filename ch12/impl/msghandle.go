package impl

import (
	"fmt"
	"go-socket/ch12/conf"
	"go-socket/ch12/iface"
)

type MsgHandle struct {
	Apis           map[uint32]iface.IRouter //map:key为msgId value为处理handle
	WorkerPoolSize uint32                   //worker的数量
	TaskQueue      []chan iface.IRequest    //每个worker对应一个消息队列
}

func NewMsgHandle() iface.IMsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]iface.IRouter),
		WorkerPoolSize: conf.WorkerPoolSize,
		TaskQueue:      make([]chan iface.IRequest, conf.WorkerPoolSize),
	}
}

func (m *MsgHandle) DoMsgHandler(request iface.IRequest) {
	handler, ok := m.Apis[request.GetMstID()]
	if !ok {
		fmt.Println("api msgId: ", request.GetMstID(), " not found!")
		return
	}

	// 执行对应的处理方法
	handler.PreHandle(request)
	handler.Handle(request)
	handler.NextHandle(request)
}

func (m *MsgHandle) AddRouter(msgId uint32, router iface.IRouter) {
	if _, ok := m.Apis[msgId]; ok {
		panic("repeated api")
	}

	m.Apis[msgId] = router
}

func (m *MsgHandle) StartWorkerPool() {
	for i := 0; i < int(conf.WorkerPoolSize); i++ {
		// 一个worker对应一个queue
		m.TaskQueue[i] = make(chan iface.IRequest, conf.TaskQueueMaxSize)
		go m.StartOneWork(i, m.TaskQueue[i])
	}
}

func (m *MsgHandle) StartOneWork(workerId int, taskQueue chan iface.IRequest) {
	fmt.Println("StartOneWork workerId: ", workerId)
	// 不断的从对队中取出消息并处理，每个worker目前不会退出
	for {
		select {
		case request := <-taskQueue:
			m.DoMsgHandler(request)
		}
	}
}

func (m *MsgHandle) SendMsgToTaskQueue(request iface.IRequest) {
	// 根据connId取模来分配队列
	workId := request.GetConn().GetConnID() % conf.WorkerPoolSize
	m.TaskQueue[workId] <- request
}
