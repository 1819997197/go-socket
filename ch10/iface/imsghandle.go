package iface

type IMsgHandle interface {
	DoMsgHandler(request IRequest)          //处理消息
	AddRouter(msgId uint32, router IRouter) //配置路由
}
