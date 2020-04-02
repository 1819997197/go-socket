package conf

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type config struct {
	WorkerPoolSize     uint32 `yaml:"worker-pool-size"`       //worker工作池最大容量
	TaskQueueMaxSize   uint32 `yaml:"task-queue-max-size"`    //消息队列允许的最大容量
	ConnLinkMaxSize    int    `yaml:"conn-link-max-size"`     //服务器允许的最大链接数
	MsgBuffChanMaxSize int    `yaml:"msg-buff-chan-max-size"` //消息管道最大允许的缓冲长度
}

var GlobalConfig *config

func init() {
	GlobalConfig = &config{
		WorkerPoolSize:     5,
		TaskQueueMaxSize:   1024,
		ConnLinkMaxSize:    2048,
		MsgBuffChanMaxSize: 512,
	}
}

func (c *config) ReadConfigFile(path string) error {
	configFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer configFile.Close()

	content, err := ioutil.ReadAll(configFile)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(content, GlobalConfig)
}
