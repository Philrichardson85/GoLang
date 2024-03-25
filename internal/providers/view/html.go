package view

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"demoBlog/pkg/converters"
	"demoBlog/pkg/sessions"
)

func WithGlobalData(c *gin.Context, data gin.H) gin.H {
	data["APP_NAME"] = viper.Get("App.Name")
	data["ERRORS"] = converters.StringToMap(sessions.Flash(c,"errors"))
	data["OLD"] = converters.StringToUrlValues(sessions.Flash(c,"old"))
	return data
}