package main

import (
	"fmt"
	"go-socket/ch14/impl"
	"io"
	"net"
	"time"
)

func main() {
	//1.创建通信socket
	conn, err := net.Dial("tcp", "127.0.0.1:8002")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("客户端通信socket创建完成!")

	for {
		dp := impl.NewDataPack()
		msg, _ := dp.Pack(impl.NewMessage(1, []byte("golang")))
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		//先读出流中的head部分
		headData := make([]byte, dp.GetHeadLen())
		n, err := io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
		if n == 0 {
			fmt.Println("服务端链接已关闭")
			break
		}
		if err != nil {
			fmt.Println("read head error")
			break
		}
		//将headData字节流 拆包到msg中
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			//msg 是有data数据的，需要再次读取data数据
			msg := msgHead.(*impl.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			//根据dataLen从io中读取字节流
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}

			fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
		}

		time.Sleep(1 * time.Second)
	}
}
