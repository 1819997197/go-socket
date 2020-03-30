package main

import (
	"fmt"
	"go-socket/ch13/iface"
	"go-socket/ch13/impl"
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
	conn.SetProperty("user_id", 1) //设置属性
	conn.SetProperty("user_name", "will")
}

func DoConnectionEnd(conn iface.IConnection) {
	fmt.Println("DoConnection end")
	if value, err := conn.GetProperty("user_id"); err == nil { //获取属性
		fmt.Println("get property user_id: ", value)
	}
	if value, err := conn.GetProperty("user_name"); err == nil {
		fmt.Println("get property user_name: ", value)
	}
}
