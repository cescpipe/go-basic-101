package main

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	LocalConfig *Local `mapstructure:"local"`
}

type Local struct {
	Server string `mapstructure:"server"`
	Port   string `mapstructure:"port"`
}

func main() {

	config := flag.String("config", "local", "function_name")
	flag.Parse()

	viper.AddConfigPath("./basic_env")
	viper.SetConfigName(fmt.Sprintf("%s", *config))
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var configVariable Config
	if err := viper.Unmarshal(&configVariable); err != nil {
		panic(fmt.Errorf("Fatal error binding config file to struct: %s \n", err))
	}
	fmt.Printf("config from yaml : %+v", viper.GetStringMap("local"))
	fmt.Println()
	fmt.Printf("config from struct : %+v", configVariable.LocalConfig)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	time.Sleep(60 * time.Second)
}
