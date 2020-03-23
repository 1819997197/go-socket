package main

import (
	"fmt"
	"go-socket/ch07/iface"
	"go-socket/ch07/impl"
)

type Router struct {
	impl.BaseRouter
}

func (r *Router) Handle(request iface.IRequest) {
	fmt.Println(request.GetConn().GetConnection().RemoteAddr().String(), "收到数据: ", string(request.GetData()))
	request.GetConn().GetConnection().Write(request.GetData())
}
