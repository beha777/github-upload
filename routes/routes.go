package routes

import (
	"github.com/gin-gonic/gin"
	"testNEW/settings"
)

func Init() {
	r := gin.Default()
	router := r.Group(settings.AppSettings.AppParams.ServerName)
	//router.GET("/sumXY", sumXY)
	router.POST("/register", userRegistration)
	router.GET("/userlist", userByName)
	router.GET("/userStartWith", userStartWith)
	r.Run(":" + settings.AppSettings.AppParams.PortRun)


}

