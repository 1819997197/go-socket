package main

import (
	"go-socket/ch14/impl"
)

func main() {
	server := impl.NewServer()

	// 设置路由
	server.AddRouter(0, &Router{})
	server.AddRouter(1, &PingRouter{})

	// 设置链接创建/断开时的Hook方法
	server.SetOnConnStart(DoConnectionBegin)
	server.SetOnConnStop(DoConnectionEnd)

	server.Run()
}
