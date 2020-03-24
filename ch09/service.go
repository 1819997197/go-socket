package main

import "go-socket/ch09/impl"

func main() {
	server := impl.NewServer()
	server.AddRouter(0, &Router{})
	server.AddRouter(1, &PingRouter{})
	server.Run()
}
