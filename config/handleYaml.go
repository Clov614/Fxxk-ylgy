package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func WriteYaml(_type interface{}, path string) {
	//dataStr为[]byte,准备写入yaml
	dataStr, err := yaml.Marshal(_type)
	if err != nil {
		log.Fatalln("WriteYaml() Error: ", err)
	}

	err = ioutil.WriteFile(path, dataStr, 0644)
	if err != nil {
		log.Fatalln("WriteYaml() writeFile Error path: "+path, err)
	}
}

func ReadYaml(_type interface{}, path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln("读取Error path: "+path, err)
	}
	err = yaml.Unmarshal(file, _type)
	if err != nil {
		log.Fatalln("ERROR:"+path+" to data error: ", err)
	}
}
