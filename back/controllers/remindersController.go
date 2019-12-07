package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) GetReminders(c *gin.Context) {
	c.String(http.StatusOK, "OK", nil)
}

func (controller *Controller) PostReminders(c *gin.Context) {
	c.String(http.StatusOK, "OK", nil)
}

func (controller *Controller) GetReminder(c *gin.Context) {
	c.String(http.StatusOK, "OK", nil)
}
