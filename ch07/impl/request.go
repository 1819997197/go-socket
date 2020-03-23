package impl

import "go-socket/ch07/iface"

type Request struct {
	Conn iface.IConnection
	Data []byte
}

func NewRequest(conn iface.IConnection, data []byte) iface.IRequest {
	return &Request{
		Conn: conn,
		Data: data,
	}
}

func (r *Request) GetConn() iface.IConnection {
	return r.Conn
}

func (r *Request) GetData() []byte {
	return r.Data
}
