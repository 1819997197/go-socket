package impl

import (
	"errors"
	"fmt"
	"go-socket/ch08/iface"
	"io"
	"net"
)

type Connection struct {
	Conn         net.Conn      //当前连接的socket TCP套接字
	ConnID       uint32        //当前连接的ID
	isClosed     bool          //当前连接的关闭状态
	Router       iface.IRouter //该连接的处理方法
	ExitBuffChan chan bool     //告知该连接已经退出/停止的channel
}

func NewConnection(conn net.Conn, connID uint32, router iface.IRouter) *Connection {
	return &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		Router:       router,
		ExitBuffChan: make(chan bool, 1),
	}
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is run")
	addr := c.RemoteAddr().String()
	for {
		dp := NewDataPack()

		// 读取客户端msg head
		headData := make([]byte, dp.GetHeadLen())
		n, err := io.ReadFull(c.GetConnection(), headData)
		if n == 0 {
			fmt.Println(addr, " 客户端已经退出")
			c.ExitBuffChan <- true
			return
		}
		if err != nil {
			fmt.Println(addr, " read msg head err: ", err)
			c.ExitBuffChan <- true
			return
		}

		// 解包，得到msgId和dataLen放到msg中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println(addr, " unpack err: ", err)
			c.ExitBuffChan <- true
			return
		}

		// 根据dataLen读取data，放在msg.Data中
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetConnection(), data); err != nil {
				fmt.Println(addr, " read msg data err: ", err)
				c.ExitBuffChan <- true
				return
			}
		}

		msg.SetData(data)

		req := NewRequest(c, msg)
		// 调用当前连接的业务方法
		go func(request iface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.NextHandle(request)
		}(req)
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

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("connection closed")
	}

	dp := NewDataPack()
	msg, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		fmt.Println("Pack err msg id: ", msgId)
		return errors.New("pack error msg")
	}

	if _, err := c.Conn.Write(msg); err != nil {
		fmt.Println("write err msg id: ", msgId)
		c.ExitBuffChan <- true
		return errors.New("conn write error")
	}

	return nil
}
