package config

import (
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		Name string `toml:"name"`
		Port int    `toml:"port"`
		Key  string `toml:"key"`
		URL  string `toml:"url"`
	} `toml:"app"`

	Database struct {
		Driver   string `toml:"driver"`
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		Name     string `toml:"name"`
	} `toml:"database"`

	Log struct {
		Driver   string `toml:"driver"`
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		Username string `toml:"username"`
		Password string `toml:"password"`
		Name     string `toml:"name"`
	} `toml:"log"`

	Mail struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		Sender   string `toml:"sender"`
		Username string `toml:"username"`
		Password string `toml:"password"`
	} `toml:"mail"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Info("No config file found, using default config")
	}

	var finalConfig AppConfig
	if err := viper.Unmarshal(&finalConfig); err != nil {
		log.Fatal("Unable to decode config file: ", err)
	}

	return &finalConfig
}
