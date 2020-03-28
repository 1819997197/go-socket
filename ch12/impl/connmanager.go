package impl

import (
	"fmt"
	"github.com/pkg/errors"
	"go-socket/ch12/iface"
	"sync"
)

type ConnManager struct {
	connections map[uint32]iface.IConnection //map: key为链接id value为链接对象
	connLock    sync.RWMutex                 //链接读写锁
}

func NewConnManager() iface.IConnManager {
	return &ConnManager{
		connections: make(map[uint32]iface.IConnection),
	}
}

func (c *ConnManager) Add(conn iface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	c.connections[conn.GetConnID()] = conn
	fmt.Println("connection add, connId: ", conn.GetConnID(), " connMapLen: ", len(c.connections))
}

func (c *ConnManager) Remove(conn iface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	delete(c.connections, conn.GetConnID())
	fmt.Println("connection remove, connId: ", conn.GetConnID(), " connMapLen: ", len(c.connections))
}

func (c *ConnManager) Get(connID uint32) (iface.IConnection, error) {
	if conn, ok := c.connections[connID]; ok {
		return conn, nil
	}

	return nil, errors.New("connection not found")
}

func (c *ConnManager) Len() int {
	return len(c.connections)
}

func (c *ConnManager) ClearConn() {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	for connID, conn := range c.connections {
		conn.Stop() //关闭链接

		delete(c.connections, connID) //清除
	}
	fmt.Println("connection clear")
}
