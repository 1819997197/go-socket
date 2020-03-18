package impl

import (
	"fmt"
	"go-socket/ch05/iface"
	"net"
)

type Server struct {
	IP       string //IP
	Port     int    //端口
	Protocol string //协议:tcp
}

func NewServer() iface.IServer {
	return &Server{
		IP:       "0.0.0.0",
		Port:     8002,
		Protocol: "tcp",
	}
}

func (s *Server) Start() {
	fmt.Printf("server start %v:%v\r\n", s.IP, s.Port)

	// 开启一个goroutine去做服务端的listen业务
	go func() {
		// 1.创建监听socket
		listener, err := net.Listen(s.Protocol, fmt.Sprintf("%s:%v", s.IP, s.Port))
		if err != nil {
			fmt.Println("net.Listen err: ", err)
			return
		}
		fmt.Println("服务端创建监听socket完成!")
		defer listener.Close()

		for {
			// 2.创建通信socket
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("listener.Accept err: ", err)
				return
			}
			fmt.Println("服务端通信socket创建完成！客户端IP:" + conn.RemoteAddr().String())

			// 目前处理比较粗暴
			go func() {
				addr := conn.RemoteAddr().String()
				buf := make([]byte, 4096)
				for {
					// 3.获取客户端发送过来的数据
					n, err := conn.Read(buf)
					if n == 0 { //客户端直接Ctrl+c退出之后，n=0&err=EOF
						fmt.Println(addr, "客户端退出!")
						return
					}
					if err != nil {
						fmt.Println(addr, "conn.Read err: ", err)
						return
					}

					fmt.Println("服务端接收到", addr, "数据为: ", string(buf[:n]))

					// 4.处理数据，返回给客户端
					data := addr + ":" + string(buf[:n])
					_, err = conn.Write([]byte(data))
					if err != nil {
						fmt.Println(addr, "conn.Write err: ", err)
						return
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("server stop")
	//TODO 资源清理
}

func (s *Server) Run() {
	s.Start()

	defer s.Stop()

	select {}
}
