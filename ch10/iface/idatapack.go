package iface

type IDataPack interface {
	GetHeadLen() uint32
	Pack(IMessage) ([]byte, error)
	Unpack([]byte) (IMessage, error)
}
