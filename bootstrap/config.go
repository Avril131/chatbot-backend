package bootstrap

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"chatbot-backend/global"
)

func InitializeConfig() *viper.Viper {
	// path to configuration file
	config := "config.yaml"

	// if have VIPER_CONFIG
	// overwrite configuration address
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}

	// init viper
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}

	// watch for changes to the configuration file
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)

		// reload
		if err := v.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})

	// load configruation insto global
	if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}

	return v
}
