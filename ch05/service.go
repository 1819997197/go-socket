package main

import "go-socket/ch05/impl"

func main() {
	server := impl.NewServer()
	server.Run()
}
