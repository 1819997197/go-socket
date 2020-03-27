package iface

type IMsgHandle interface {
	DoMsgHandler(request IRequest)          //处理消息
	AddRouter(msgId uint32, router IRouter) //配置路由
	StartWorkerPool()                       //启动worker工作池
	SendMsgToTaskQueue(request IRequest)    //将消息写入队列，交由worker处理
}
