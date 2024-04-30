package controllers

import (
	"log"
	"my-tools/thanksfulness/notion"
	"my-tools/thanksfulness/slack"

	"github.com/gin-gonic/gin"
)

func Thanksfulness(c *gin.Context) {
	thanksList, err := notion.GetThanksList()
	if err != nil {
		log.Printf("Error during Notion API: %v\n", err)
		c.AbortWithStatusJSON(500, gin.H{"message": "Notion API Failed!"})
	}

	err = slack.PostThanksList(thanksList)
	if err != nil {
		log.Printf("Error during Slack API: %v\n", err)
		c.AbortWithStatusJSON(500, gin.H{"message": "Slack API Failed!"})
	}

	c.JSON(200, gin.H{
		"message": "thanksfulness!",
	})
}
