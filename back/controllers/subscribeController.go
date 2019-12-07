package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) Subscribe(c *gin.Context) {
	c.String(http.StatusOK, "OK", nil)
}
