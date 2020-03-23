package main

import "go-socket/ch07/impl"

func main() {
	server := impl.NewServer()
	server.AddRouter(&Router{})
	server.Run()
}
