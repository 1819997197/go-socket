package main

import (
	"net"
	"fmt"
)

func main() {
	//1.创建通信socket
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("客户端通信socket创建完成!")

	//2.向服务器端传输数据
	_, err = conn.Write([]byte("hello world"))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return
	}

	buf := make([]byte, 4096)
	//3.接受服务器端返回的数据
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	//显示数据
	fmt.Println(string(buf[:n]))
}
