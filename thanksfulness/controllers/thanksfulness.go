package controllers

import "github.com/gin-gonic/gin"

func Thanksfulness(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "thanksfulness!",
	})
}
