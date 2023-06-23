package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		return err // 读取配置信息失败
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(c fsnotify.Event) {
		fmt.Println("监测配置文件更新~")
	})

	return
}
