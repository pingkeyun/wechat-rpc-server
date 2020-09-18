package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var cfg *viper.Viper

func Setup() {
	cfg = viper.New()

	cfg.AddConfigPath("configs")
	if err := cfg.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic(fmt.Errorf("Config file not found; ignore error if desired: %s \n", err))
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("%s\n", err))
		}
	}

	cfg.WatchConfig()
	cfg.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	// app
	AppConfig = NewApp(cfg.GetString("app.APPID"), cfg.GetString("app.APPSECRET"))

	// database
	DbConfig = &Database{
		Host:        cfg.GetString("database.Host"),
		User:        cfg.GetString("database.User"),
		Password:    cfg.GetString("database.Password"),
		Name:        cfg.GetString("database.name"),
		TablePrefix: cfg.GetString("database.TablePrefix"),
	}

	// redis
	RedisConfig = &Redis{
		Host:        cfg.GetString("redis.Host"),
		Password:    cfg.GetString("redis.Password"),
		MaxIdle:     cfg.GetInt("redis.maxIdle"),
		MaxActive:   cfg.GetInt("redis.MaxActive"),
		IdleTimeout: cfg.GetDuration("redis.IdleTimeout"),
	}

}
