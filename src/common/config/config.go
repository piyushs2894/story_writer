package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"story_writer/src/constant"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	gcfg "gopkg.in/gcfg.v1"
)

var CF *Config

type Config struct {
	Database map[string]*DatabaseConfig
}

type DatabaseConfig struct {
	Master        string
	Slave         string
	MaxMasterConn int
	MaxSlaveConn  int
}

func Init() {
	CF = &Config{}
	ok := ReadConfig(CF, "./files/etc", "story_writer")
	if !ok {
		log.Fatal("Failed to read config file")
	}
}

// ReadConfig is file handler for reading configuration files into variable
// Param: - config pointer of Config, filepath string
// Return: - boolean
func ReadConfig(cfg *Config, path string, module string) bool {
	environ := os.Getenv("USERENV")
	if environ == "" {
		environ = constant.ENV_DEVELOPMENT
	}

	environ = strings.ToLower(environ)

	parts := []string{"main"}
	var configString []string

	for _, v := range parts {
		fname := path + "/" + module + "/" + environ + "/" + module + "." + v + ".ini"
		fmt.Println(time.Now().Format("2006/01/02 15:04:05"), "Reading", fname)

		config, err := ioutil.ReadFile(fname)
		if err != nil {
			log.Errorln("common/config.go function ReadConfig", err)
			return false
		}

		configString = append(configString, string(config))
	}

	err := gcfg.ReadStringInto(cfg, strings.Join(configString, "\n\n"))

	if err != nil {
		log.Errorln("common/config.go function ReadConfig", err)
		return false
	}

	return true
}

func GetConfig() *Config {
	return CF
}
