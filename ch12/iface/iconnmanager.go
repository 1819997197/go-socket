package iface

type IConnManager interface {
	Add(conn IConnection)                   //添加链接
	Remove(conn IConnection)                //移出链接
	Get(connID uint32) (IConnection, error) //根据链接id获取链接
	Len() int                               //获取当前链接数
	ClearConn()                             //清除并关闭所有链接
}
