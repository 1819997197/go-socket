package impl

import (
	"errors"
	"fmt"
	"go-socket/ch12/conf"
	"go-socket/ch12/iface"
	"io"
	"net"
)

type Connection struct {
	TcpServer    iface.IServer
	Conn         net.Conn         //当前连接的socket TCP套接字
	ConnID       uint32           //当前连接的ID
	isClosed     bool             //当前连接的关闭状态
	MsgHandle    iface.IMsgHandle //该连接的处理方法
	ExitBuffChan chan bool        //告知该连接已经退出/停止的channel
	msgChan      chan []byte      //业务方法处理完的数据, 读写分离管道(无缓冲管道)
	msgBuffChan  chan []byte      //业务方法处理完的数据, 读写分离管道(有缓冲管道)
}

func NewConnection(server iface.IServer, conn net.Conn, connID uint32, msgHandle iface.IMsgHandle) iface.IConnection {
	connection := &Connection{
		TcpServer:    server,
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		MsgHandle:    msgHandle,
		ExitBuffChan: make(chan bool, 1),
		msgChan:      make(chan []byte),
		msgBuffChan:  make(chan []byte, conf.MsgBuffChanMaxSize),
	}

	connection.TcpServer.GetConnManager().Add(connection) //将新增的链接添加到链接管理map中
	return connection
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
		c.MsgHandle.SendMsgToTaskQueue(req) //把消息直接写入队列
	}
}

func (c *Connection) StartWrite() {
	fmt.Println("Write Goroutine is run")
	for {
		select {
		case msg := <-c.msgChan:
			if _, err := c.Conn.Write(msg); err != nil {
				fmt.Println("write err msg info: ", msg)
				c.ExitBuffChan <- true
			}
		case msg, ok := <-c.msgBuffChan:
			if ok {
				if _, err := c.Conn.Write(msg); err != nil {
					fmt.Println("write err msg info: ", msg)
					c.ExitBuffChan <- true
				}
			}
		case <-c.ExitBuffChan: //退出
			c.Stop()
			return
		}
	}
}

func (c *Connection) Start() {
	go c.StartReader()

	go c.StartWrite()

	c.TcpServer.CallOnConnStart(c) //链接创建时Hook方法
}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}

	c.isClosed = true

	c.TcpServer.GetConnManager().Remove(c) //链接管理map中移出链接
	c.TcpServer.CallOnConnStop(c)          //链接断开时Hook方法

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

	c.msgChan <- msg

	return nil
}

func (c *Connection) SendBuffMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New("connection closed")
	}

	dp := NewDataPack()
	msg, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		fmt.Println("Pack err msg id: ", msgId)
		return errors.New("pack error msg")
	}

	c.msgBuffChan <- msg

	return nil
}
