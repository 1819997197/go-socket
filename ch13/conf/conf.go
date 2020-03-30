package conf

// 暂时先直接定义变量，后续给定默认值并从配置文件读取
var WorkerPoolSize uint32 = 5 //worker工作池最大容量

var TaskQueueMaxSize uint32 = 1024 //消息队列允许的最大容量

var ConnLinkMaxSize int = 2048 //服务器允许的最大链接数

var MsgBuffChanMaxSize int = 512 //消息管道最大允许的缓冲长度
