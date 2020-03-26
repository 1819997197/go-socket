package iface

import "net"

type IConnection interface {
	Start()                                  //启动连接
	Stop()                                   //结束当前连接状态
	GetConnection() net.Conn                 //从当前连接获取原始的socket conn
	GetConnID() uint32                       //获取当前连接ID
	RemoteAddr() net.Addr                    //获取远程客户端地址信息
	SendMsg(msgId uint32, data []byte) error //将message发送给远程的TCP客户端
}
