package main

import "go-socket/ch08/impl"

func main() {
	server := impl.NewServer()
	server.AddRouter(&Router{})
	server.Run()
}
