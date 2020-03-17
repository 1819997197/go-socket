# socket编程

## 1.并发服务端
```
TCP-CS并发服务器：

	1.  创建 监听套接字 listener := net.Listen("tcp"， 服务器的IP+port)		// tcp 不能大写

	2.  defer listener.Close()

	3.  for 循环 阻塞监听 客户端连接事件 	conn := listener.Accept()

	4. 创建 go程 对应每一个 客户端进行数据通信	go HandlerConnet()

	5. 实现 handlerConn(conn net.Conn)

		1) defer conn.Close()

		2) 获取成功连接的客户端 Addr 		conn.RemoteAddr()

		3) for 循环 读取 客户端发送数据		conn.Read(buf)

		4) 处理数据 小 —— 大	strings.ToUpper()

		5）回写转化后的数据		conn.Write(buf[:n])
```

## 2.客户端
```
TCP-CS并发客户端：

	1. 匿名 go 程 ， 获取 键盘输入， 写给服务器

	2. for 循环读取服务器回发数据

	发送数据时，默认在结尾自带换行键
```

## 3.Usage

Run the service
```
go run service.go
```

Run the client
```
// 可同时运行多个服务端
go run client.go
```


服务端可并发接收多个客户端请求，所有处理逻辑放在一个文件里面，不利于后续的扩展，下一节将会吧服务端面向interface来实现。