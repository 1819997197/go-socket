package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//1.创建监听socket
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()
	fmt.Println("监听socket创建完成!")

	//2.创建通信socket(阻塞等待用户连接)
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("通信socket创建完成!")

	buf := make([]byte, 4096)
	//3.读客户端写入的数据
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	fmt.Println("接收到服务端数据为:", string(buf[:n]))

	//4.数据处理之后，回传给客户端
	conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
}
