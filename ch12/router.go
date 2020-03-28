package main

import (
	"fmt"
	"go-socket/ch12/iface"
	"go-socket/ch12/impl"
)

type Router struct {
	impl.BaseRouter
}

func (r *Router) Handle(request iface.IRequest) {
	//fmt.Println(request.GetConn().GetConnection().RemoteAddr().String(), "收到数据: ", string(request.GetData()))
	request.GetConn().SendMsg(request.GetMstID(), request.GetData())
}

type PingRouter struct {
	impl.BaseRouter
}

func (r *PingRouter) Handle(request iface.IRequest) {
	//fmt.Println(request.GetConn().GetConnection().RemoteAddr().String(), "收到数据: ", string(request.GetData()))
	request.GetConn().SendBuffMsg(request.GetMstID(), request.GetData())
}

func DoConnectionBegin(conn iface.IConnection) {
	fmt.Println("DoConnection begin")
	conn.SendMsg(3, []byte("DoConnection begin"))
}

func DoConnectionEnd(conn iface.IConnection) {
	fmt.Println("DoConnection end")
}
