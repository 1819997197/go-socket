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
客户端请求的链接，目前是绑死了(handleAPI方法)处理业务，下一节将提供用户可自定义的路由处理业务方法