package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"strings"
	"time"
)

type Configuration struct {
	App App `yaml:"app"`
	Db  Db  `yaml:"db"`
}

type App struct {
	Name    string `yaml:"name"`
	Profile string `yaml:"profile"`
	Timeout int    `yaml:"timeout"`
	Port    int    `yaml:"port"`
}

type Db struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbName"`
	Timezone string `yaml:"timezone"`
	Timeout  int    `yaml:"timeout"`
}

func LoadConfig() (config *Configuration, e error) {
	initial := time.Now()

	logrus.Infof("[Env Config] Initializing env variable configurations")

	viper.SetConfigName("configs/application")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Load Config from File
	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			logrus.Warnf("No Config file found, loaded config from Environment - Default path ./conf")
		default:
			return nil, errors.Wrap(err, "config.LoadConfig")
		}
	}

	err := viper.Unmarshal(&config)

	if err != nil {
		logrus.Error("[Env Config] Error serializing config", err)
		return nil, errors.Wrap(err, "[Env Config] Error serializing config")
	}

	delta := time.Since(initial).Milliseconds()
	logrus.Infof(fmt.Sprintf("[Env Config] Finalized env variable configurations in %dus", delta))

	return config, nil
}
