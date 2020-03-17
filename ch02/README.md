# socket编程简介

## 1.什么是socket？
```
socket的原意是“插座”, 在计算机通信领域, socket被翻译为“套接字”, 它是计算机之间进行通信的一种约定或一种方式。
```

## 2.socket的本质是什么？
```
socket的诞生是为了应用程序能够更方便的将数据经由传输层来传输。
socket本质是编程接口(API), 对TCP/IP的封装，然后应用层程序直接调用socket API进行通信
```

## 3.查看实际的套接字？
```
套接字的实体就是通信控制信息，包含了本地主机IP+端口、远程主机IP+端口、TCP/UDP。
用netstat -ano命令显示套接字内容，每一行就相当于一个套接字。
```

## 4.socket在网络中所处的位置
![img/socket-01.jpg](https://github.com/1819997197/go-socket/blob/master/ch02/img/socket-01.jpg)

## 5.socket通信模型

#### (1)建立TCP/IP连接的过程(三次握手)
![img/socket-02.jpg](https://github.com/1819997197/go-socket/blob/master/ch02/img/socket-02.jpg)

#### (2)交换数据的过程
![img/socket-03.jpg](https://github.com/1819997197/go-socket/blob/master/ch02/img/socket-03.jpg)

#### (3)TCP/IP连接关闭的过程(四次挥手)
![img/socket-04.jpg](https://github.com/1819997197/go-socket/blob/master/ch02/img/socket-04.jpg)

## TODO
这一节只对socket作了一些介绍，下一节将进行实际的代码编写