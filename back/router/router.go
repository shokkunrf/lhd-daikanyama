package router

import (
	"database/sql"
	"lhd-daikanyama/controllers"
	"os"

	"github.com/gin-gonic/gin"
)

func GetRoute() (*gin.Engine, error) {
	r := gin.Default()
	ctrl := controllers.GetController()

	user := os.Getenv("MYSQL_USER")
	ps := os.Getenv("MYSQL_PASSWORD")
	protocol := "tcp"
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DATABASE")
	dbarg := "?charset=utf8&parseTime=true"
	db, err := sql.Open("mysql", user+":"+ps+"@"+protocol+"("+host+":"+port+")/"+dbname+dbarg)
	if err != nil {
		return nil, err
	}

	ctrl.DB = db

	r.GET("/", ctrl.GetRoot)
	r.GET("/reminders", ctrl.GetReminders)
	r.POST("/reminders", ctrl.PostReminders)
	r.GET("/reminders/:id", ctrl.GetReminder)
	r.POST("/subscribe", ctrl.Subscribe)
	r.POST("/signup", ctrl.PostSignup)
	r.POST("/signin", ctrl.PostSignin)

	return r, err
}
