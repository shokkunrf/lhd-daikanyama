package router

import (
	"lhd-daikanyama/controllers"

	"github.com/gin-gonic/gin"
)

func GetRoute() *gin.Engine {
	r := gin.Default()
	ctrl := controllers.GetController()

	r.GET("/", ctrl.GetRoot)
	r.GET("/reminders", ctrl.GetReminders)
	r.POST("/reminders", ctrl.PostReminders)
	r.GET("/reminders/:id", ctrl.GetReminder)
	r.POST("/subscribe", ctrl.Subscribe)
	r.GET("/signin", ctrl.GetLogin)
	r.POST("/signup", ctrl.PostSignup)
	r.POST("/signin", ctrl.PostSignin)

	return r
}
