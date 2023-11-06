package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health_check() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Checking Health": "the server is up and running",
		})
	}
}
