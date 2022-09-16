package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Token      string `yaml:"token"`       // t
	CostTime   int    `yaml:"cost_time"`   //完成耗时
	CycleCount int    `yaml:"cycle_count"` //通关次数

}

var (
	config    Config
	UserAgent = "Mozilla/5.0 (iPad; CPU OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.27(0x18001b36) NetType/WIFI Language/zh_CN"
	helpStr   = "#token为获取到的header中t值\n#cost_time为完成耗时 单位s，默认-1随机表示随机生成1s~1h之内的随机数，设置为正数则为固定\n#需要通关的次数，最大支持20，默认1\n"
)

func init() {
	if !PathExists("./config.yml") {
		file, err := os.OpenFile("./config.yml", os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatalln("WriteYaml() writeFile Error path: "+"./config.yml", err)
		}
		file.Write([]byte(helpStr))
		config.CycleCount = 1
		config.CostTime = -1
		// 写入config
		configByte, _ := yaml.Marshal(config)
		file.Write(configByte)

		log.Error("初始化成功，请重新启动")
		os.Exit(0)
	}
}

func ReadConfig() *Config {
	ReadYaml(&config, "./config.yml")
	return &config
}
