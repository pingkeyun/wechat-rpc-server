package config

import "sync"

var AppConfig App

type App struct {
	AppId, AppSecret string
}

func NewApp(appId, secret string) App {
	var once sync.Once
	once.Do(func() {
		AppConfig = App{
			AppId:     appId,
			AppSecret: secret,
		}
	})

	return AppConfig
}

// GetAppId
func (a App) GetAppId() string {
	return a.AppId
}

// GetSecret
func (a App) GetSecret() string {
	return a.AppSecret
}
