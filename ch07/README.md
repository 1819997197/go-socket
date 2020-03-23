# socket编程

## 1.Usage

Run the service
```
go run service.go
```

Run the client
```
// 可同时运行多个服务端
go run client.go
```

## TODO
实现了一个简单的路由配置。现在我们把服务器的全部数据放到一个Request里面，用一个[]byte来接收全部的数据，既没有长度，也没有类型。
下一节将封装一种消息类型，来存放服务器数据。