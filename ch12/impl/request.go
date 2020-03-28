package impl

import "go-socket/ch12/iface"

type Request struct {
	Conn iface.IConnection
	Msg  iface.IMessage
}

func NewRequest(conn iface.IConnection, msg iface.IMessage) iface.IRequest {
	return &Request{
		Conn: conn,
		Msg:  msg,
	}
}

func (r *Request) GetConn() iface.IConnection {
	return r.Conn
}

func (r *Request) GetData() []byte {
	return r.Msg.GetData()
}

func (r *Request) GetMstID() uint32 {
	return r.Msg.GetMsgID()
}
