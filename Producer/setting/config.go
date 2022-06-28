package setting

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	KafkaConfig `mapstructure:"kafka"`
}
type KafkaConfig struct {
	Address  string `mapstructure:"address"`
	ChanSize int64  `mapstructure:"chan_size"`
}

var conf = new(Config)

func Init() (err error) {
	viper.SetConfigFile("./config/config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := viper.Unmarshal(conf); err != nil {
		fmt.Println(err)
		return err
	}
	return
}
func GetConf() *Config {
	return conf
}
