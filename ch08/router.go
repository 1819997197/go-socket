package main

import (
	"fmt"
	"go-socket/ch08/iface"
	"go-socket/ch08/impl"
)

type Router struct {
	impl.BaseRouter
}

func (r *Router) Handle(request iface.IRequest) {
	fmt.Println(request.GetConn().GetConnection().RemoteAddr().String(), "收到数据: ", string(request.GetData()))
	request.GetConn().SendMsg(1, request.GetData())
}
