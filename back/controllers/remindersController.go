package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type resposeGetReminders struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	UserName string    `json:"user_name"`
	Date     time.Time `json:"date"`
}

type responseGetReminder struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	UserName string `json:"user_name"`
}

func (controller *Controller) GetReminders(c *gin.Context) {
	c.String(http.StatusOK, "OK", nil)
}

func (controller *Controller) PostReminders(c *gin.Context) {
	c.String(http.StatusOK, "OK", nil)
}

func (controller *Controller) GetReminder(c *gin.Context) {
	c.String(http.StatusOK, "OK", nil)
}
