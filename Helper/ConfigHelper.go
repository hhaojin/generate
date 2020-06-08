package Helper

import (
	"fmt"
	"github.com/go-yaml/yaml"
)

//加载yaml 配置文件
type SysConfig struct {
	DB struct {
		Driver string
		DSN    string
	}
}

func LoadConfig() (*SysConfig, error) {
	conf_path := WorkDir + SYS_CONFIG_PATH
	if !IsFileExist(conf_path) {
		return nil, fmt.Errorf("missing config file!")
	}
	if config_bytes := ReadFile(conf_path); config_bytes == nil {
		return nil, fmt.Errorf("load config error!")
	} else {
		config := &SysConfig{}
		err := yaml.Unmarshal(config_bytes, config)
		if err != nil {
			return nil, fmt.Errorf("config parse-error:%s", err.Error())
		}
		return config, nil
	}

}
