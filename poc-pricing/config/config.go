package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/hasemeneh/poc-material/poc-pricing/pkg/logger"
)

type config struct {
	readfile readFileFunc
	fatalln  logger.LogFunc
}
type ConfigManager interface {
	Read() *MainConfig
}
type readFileFunc func(filename string) ([]byte, error)

func New() ConfigManager {
	return &config{
		readfile: ioutil.ReadFile,
		fatalln:  log.Fatalln,
	}
}

func (c *config) Read() *MainConfig {
	cfg := MainConfig{}
	configByte, err := c.readfile("/etc/hasemeneh/app/configs/config.json")
	if err != nil {
		c.fatalln("Something Wrong Happen", err.Error())
	}
	err = json.Unmarshal(configByte, &cfg)
	if err != nil {
		c.fatalln("Something Wrong Happen", err.Error())
	}
	return &cfg
}

// using driver yaml
type Config struct {
	Server ServerConfig `yaml:"Server"`
	DB     DBConfig     `yaml:"DB"`
}

type (
	ServerConfig struct {
		Port                    string `yaml:"Port"`
		BasePath                string `yaml:"BasePath"`
		GracefulTimeoutInSecond int    `yaml:"GracefulTimeout"`
		ReadTimeoutInSecond     int    `yaml:"ReadTimeout"`
		WriteTimeoutInSecond    int    `yaml:"WriteTimeout"`
		APITimeout              int    `yaml:"APITimeout"`
	}
	DBConfig struct {
		RetryInterval int    `yaml:"RetryInterval"`
		MaxIdleConn   int    `yaml:"MaxIdleConn"`
		MaxConn       int    `yaml:"MaxConn"`
		DSN           string `yaml:"DSN"`
	}
)
