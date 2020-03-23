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
实现了消息封装，message包含消息ID、数据长度、数据三个成员；数据传输协议：head + body，head包含数据长度和消息id各占4个字节；
下一节实现多路由配置