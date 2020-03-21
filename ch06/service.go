package main

import "go-socket/ch06/impl"

func main() {
	server := impl.NewServer()
	server.Run()
}
