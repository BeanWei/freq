package g

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var (
	cfg         *conf
	cfgOnce     sync.Once
	defaultConf = []byte(`
server:
  address: localhost:8888
`)
)

type conf struct {
	Server struct {
		Address string `yaml:"address"`
	} `yaml:"server"`
}

func Cfg() *conf {
	cfgOnce.Do(func() {
		viper.SetConfigType("yaml")
		viper.AutomaticEnv()

		cfgfile := viper.GetString("cfgfile")
		if cfgfile == "" {
			cfgfile = "../../config/config.yaml"
		}
		viper.SetConfigFile(cfgfile)

		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file: ", viper.ConfigFileUsed())
		} else if err := viper.ReadConfig(bytes.NewBuffer(defaultConf)); err != nil {
			panic(fmt.Errorf("read config faild: %+v", err))
		}

		if err := viper.Unmarshal(&cfg); err != nil {
			panic(fmt.Errorf("viper unmarshal config faild: %+v", err))
		}
	})
	return cfg
}
