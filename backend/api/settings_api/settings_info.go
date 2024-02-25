package settings_api

import (
	"github.com/576690/qy8/backend/model/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.FailWithCode(res.SettingsError, c)
}
