package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) GetLogin(c *gin.Context) {
	c.String(http.StatusOK, "OK", nil)
}
func (controller *Controller) PostSignup(c *gin.Context) {
	c.String(http.StatusOK, "OK", nil)
}

func (controller *Controller) PostSignin(c *gin.Context) {
	c.String(http.StatusOK, "OK", nil)
}
