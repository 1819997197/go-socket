# socket编程

## 1.原理
![img/socket-01.png](https://github.com/1819997197/go-socket/blob/master/ch03/img/socket-01.png)

## 2.服务端
```
TCP-CS服务器：

	1.  创建监听socket  listener := net.Listen("TCP", "IP+port")	IP+port	—— 服务器自己的IP 和 port

	2.  启动监听  conn := listener.Accept()  conn 用于 通信的 socket

	3.  conn.Read()

	4.  处理使用 数据

	5.  conn.Write()

	6.  关闭  listener、conn
```

## 3.客户端
```
TCP-CS客户端：

	1.  conn, err := net.Dial("TCP", 服务器的IP+port)

	2.  写数据给 服务器 conn.Write()

	3.  读取服务器回发的 数据 conn.Read()

	4.  conn.Close()
```

## 4.Usage

Run the service
```
go run service.go
```

Run the client
```
go run client.go
```

## TODO
服务端只运行了一次，就结束了进程。显然达不到要求，下一节中，将会调整为并发服务端，可接收多个客户端请求