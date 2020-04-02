package iface

type IRequest interface {
	GetConn() IConnection
	GetData() []byte
	GetMstID() uint32
}
