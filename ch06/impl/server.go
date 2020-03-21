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

		// 每个连接有一个id
		var cid uint32
		cid = 0

		for {
			// 2.创建通信socket
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("listener.Accept err: ", err)
				return
			}
			fmt.Println("服务端通信socket创建完成！客户端IP:" + conn.RemoteAddr().String())

			// 处理连接请求的业务
			dealConn := NewConnection(conn, cid, handleAPI)
			cid++
			go dealConn.Start()
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

func handleAPI(conn net.Conn, data []byte, n int) error {
	fmt.Println(conn.RemoteAddr().String(), "收到数据: ", string(data[:n]))
	_, err := conn.Write(data[:n])
	return err
}
