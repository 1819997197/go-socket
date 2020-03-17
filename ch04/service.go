package main

import (
	"fmt"
	"net"
	"strings"
)

//通信socket处理客户端发送过来的数据
func handlerConn(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	buf := make([]byte, 4096)
	for {
		//1.读取客户端发送过来的数据
		n, err := conn.Read(buf)
		if string(buf[:n]) == "exit\r\n" || string(buf[:n]) == "exit\r" || string(buf[:n]) == "exit\n" {
			fmt.Println(addr + "客户端退出!")
			return
		}
		if n == 0 {
			fmt.Println(addr + "客户端退出!!")
			return
		}
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}

		fmt.Printf("服务端接收到%v数据%v\r\n", addr, string(buf[:n]))
		//2.处理完数据发送回客户端
		str := strings.ToUpper(addr + ": " + string(buf[:n]))
		conn.Write([]byte(str))
	}
}

func main() {
	//1.创建监听socket
	listener, err := net.Listen("tcp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()
	fmt.Println("服务端创建监听socket完成!")

	for {
		//2.创建通信socket
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}
		fmt.Println("服务端通信socket创建完成! 客户端IP:" + conn.RemoteAddr().String())
		go handlerConn(conn)
	}
}
