package iface

type IServer interface {
	Start()                                 //服务初始化
	Stop()                                  //服务停止后，资源清理
	Run()                                   //运行服务
	AddRouter(msgId uint32, router IRouter) //设置路由
	GetConnManager() IConnManager           //获取链接管理
	SetOnConnStart(func(IConnection))       //设置链接创建时的Hook函数
	SetOnConnStop(func(IConnection))        //设置链接断开时的Hook函数
	CallOnConnStart(IConnection)            //调用连接SetOnConnStart Hook函数
	CallOnConnStop(IConnection)             //调用连接SetOnConnStop Hook函数
}
