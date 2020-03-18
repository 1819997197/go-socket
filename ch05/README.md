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
一个基础的server框架已经有了基本的雏形。服务端现在针对具体的客户端链接只作了简单的数据拼接功能(ip+port+data)，下一节将连接抽象出来，绑定到具体的业务方法上