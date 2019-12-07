package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func GetController() *Controller {
	controller := Controller{}
	return &controller
}

func (controller *Controller) GetRoot(c *gin.Context) {
	c.String(http.StatusOK, "OK", nil)
}
