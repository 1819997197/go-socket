package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//1.创建通信socket
	conn, err := net.Dial("tcp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("客户端通信socket创建完成!")

	//2.不断获取键盘输入，并发送给客户端
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				fmt.Println("os.Stdin.Read err:", err)
				return
			}

			//3.向服务器端传输数据
			conn.Write(buf[:n])
		}
	}()

	for {
		//4.接收服务端返回的数据
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("服务端连接已关闭!")
			return
		}
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}

		fmt.Println(string(buf[:n]))
	}
}
