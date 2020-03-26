package iface

type IRouter interface {
	PreHandle(request IRequest)  //业务处理方法之前执行
	Handle(request IRequest)     //业务处理方法
	NextHandle(request IRequest) //业务处理方法之后执行
}
