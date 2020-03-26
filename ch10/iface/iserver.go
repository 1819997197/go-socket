package iface

type IServer interface {
	Start()                                 //服务初始化
	Stop()                                  //服务停止后，资源清理
	Run()                                   //运行服务
	AddRouter(msgId uint32, router IRouter) //设置路由
}
