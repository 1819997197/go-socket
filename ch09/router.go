package main

import (
	"fmt"
	"go-socket/ch09/iface"
	"go-socket/ch09/impl"
)

type Router struct {
	impl.BaseRouter
}

func (r *Router) Handle(request iface.IRequest) {
	fmt.Println(request.GetConn().GetConnection().RemoteAddr().String(), "收到数据: ", string(request.GetData()))
	request.GetConn().SendMsg(request.GetMstID(), request.GetData())
}

type PingRouter struct {
	impl.BaseRouter
}

func (r *PingRouter) Handle(request iface.IRequest) {
	fmt.Println(request.GetConn().GetConnection().RemoteAddr().String(), "收到数据: ", string(request.GetData()))
	request.GetConn().SendMsg(request.GetMstID(), request.GetData())
}
