package impl

import "go-socket/ch09/iface"

type BaseRouter struct{}

// BaseRouter空实现，用户自定义路由，继承这个结构体，就不用实现iface.IRequest所有接口
func (router *BaseRouter) PreHandle(request iface.IRequest) {}

func (router *BaseRouter) Handle(request iface.IRequest) {}

func (router *BaseRouter) NextHandle(request iface.IRequest) {}
