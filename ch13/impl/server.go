package impl

import (
	"fmt"
	"go-socket/ch13/conf"
	"go-socket/ch13/iface"
	"net"
)

type Server struct {
	IP                string //IP
	Port              int    //端口
	Protocol          string //协议:tcp
	MsgHandle         iface.IMsgHandle
	ConnectionManager iface.IConnManager

	// 链接启动/停止自定义Hook方法
	OnConnStart func(connection iface.IConnection)
	OnConnStop  func(connection iface.IConnection)
}

func NewServer() iface.IServer {
	return &Server{
		IP:                "0.0.0.0",
		Port:              8002,
		Protocol:          "tcp",
		MsgHandle:         NewMsgHandle(),
		ConnectionManager: NewConnManager(),
	}
}

func (s *Server) Start() {
	fmt.Printf("server start %v:%v\r\n", s.IP, s.Port)

	// 开启一个goroutine去做服务端的listen业务
	go func() {
		// 开启worker工作池
		s.MsgHandle.StartWorkerPool()

		// 创建监听socket
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
			// 创建通信socket
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("listener.Accept err: ", err)
				return
			}
			fmt.Println("服务端通信socket创建完成！客户端IP:" + conn.RemoteAddr().String())

			// 判断链接数是否超出设置
			if s.ConnectionManager.Len() > conf.ConnLinkMaxSize-1 { //先判断长度，后加入到map中，所以需要减一，最终最大链接数为conf.ConnLinkMaxSize个
				conn.Close()
				continue
			}

			// 处理连接请求的业务
			dealConn := NewConnection(s, conn, cid, s.MsgHandle)
			cid++
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("server stop")
	//TODO 资源清理

	s.ConnectionManager.ClearConn() //服务关闭时，清除所有的链接信息
}

func (s *Server) Run() {
	s.Start()

	defer s.Stop()

	select {}
}

func (s *Server) AddRouter(msgId uint32, router iface.IRouter) {
	s.MsgHandle.AddRouter(msgId, router)
}

func (s *Server) GetConnManager() iface.IConnManager {
	return s.ConnectionManager
}

func (s *Server) SetOnConnStart(hookFunc func(iface.IConnection)) {
	s.OnConnStart = hookFunc
}

func (s *Server) SetOnConnStop(hookFunc func(iface.IConnection)) {
	s.OnConnStop = hookFunc
}

func (s *Server) CallOnConnStart(conn iface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("CallOnConnStart")
		s.OnConnStart(conn)
	}
}

func (s *Server) CallOnConnStop(conn iface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("CallOnConnStop")
		s.OnConnStop(conn)
	}
}
