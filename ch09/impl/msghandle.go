package impl

import (
	"fmt"
	"go-socket/ch09/iface"
)

type MsgHandle struct {
	Apis map[uint32]iface.IRouter //map:key为msgId value为处理handle
}

func NewMsgHandle() iface.IMsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]iface.IRouter),
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
