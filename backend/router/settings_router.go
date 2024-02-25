package router

import (
	"github.com/576690/qy8/backend/api"
)

func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings", settingsApi.SettingsInfoView)
}
