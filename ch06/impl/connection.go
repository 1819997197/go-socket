package impl

import (
	"fmt"
	"go-socket/ch06/iface"
	"net"
)

type Connection struct {
	Conn         net.Conn       //当前连接的socket TCP套接字
	ConnID       uint32         //当前连接的ID
	isClosed     bool           //当前连接的关闭状态
	handleAPI    iface.HandFunc //该连接的处理方法api
	ExitBuffChan chan bool      //告知该连接已经退出/停止的channel
}

func NewConnection(conn net.Conn, connID uint32, handleAPI iface.HandFunc) *Connection {
	return &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		handleAPI:    handleAPI,
		ExitBuffChan: make(chan bool, 1),
	}
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is run")
	addr := c.RemoteAddr().String()
	for {
		buf := make([]byte, 512)
		n, err := c.Conn.Read(buf)
		if n == 0 {
			fmt.Println(addr, "客户端已退出")
			c.ExitBuffChan <- true
			return
		}
		if err != nil {
			fmt.Println(addr, "conn.Read err: ", err)
			c.ExitBuffChan <- true
			return
		}

		// 调用当前连接的业务方法
		if err := c.handleAPI(c.Conn, buf, n); err != nil {
			fmt.Println("connID ", c.ConnID, " handle id err: ", err)
			c.ExitBuffChan <- true
			return
		}
	}
}

func (c *Connection) Start() {
	go c.StartReader()

	for {
		select {
		case <-c.ExitBuffChan: //退出
			c.Stop()
			return
		}
	}
}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}

	c.isClosed = true

	// 关闭socket链接
	c.Conn.Close()

	// 关闭该连接全部管道
	close(c.ExitBuffChan)
}

func (c *Connection) GetConnection() net.Conn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
