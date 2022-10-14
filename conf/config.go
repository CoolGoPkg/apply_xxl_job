package conf

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type DemoConfig struct {
	XXLJobConf *XXLJobConfig `yaml:"xxl_job_conf"`
}

type XXLJobConfig struct {
	Token        string `yaml:"token"`
	AppName      string `yaml:"app_name"`
	ClientPort   int    `yaml:"client_port"`
	AdminAddress string `yaml:"admin_address"`
}

var Config = &DemoConfig{}

func load(config interface{}, file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	switch {
	case strings.HasSuffix(file, ".yaml") || strings.HasSuffix(file, ".yml"):
		return yaml.Unmarshal(data, config)
	case strings.HasSuffix(file, ".json"):
		return json.Unmarshal(data, config)
	default:
		if json.Unmarshal(data, config) != nil {
			if yaml.Unmarshal(data, config) != nil {
				return errors.New("failed to decode config")
			}
		}
		return nil
	}
}

func InitConf(file string) {
	err := load(Config, file)
	if err != nil {
		panic(err)
	}
	fmt.Println(Config)
}
