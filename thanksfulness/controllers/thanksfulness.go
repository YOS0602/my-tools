package controllers

import (
	"log"
	"my-tools/thanksfulness/notion"

	"github.com/gin-gonic/gin"
	"github.com/slack-go/slack"
)

func Thanksfulness(c *gin.Context) {
	thanksList, err := notion.GetThanksLists()
	if err != nil {
		log.Printf("Error during Notion API: %v\n", err)
		c.AbortWithStatusJSON(500, gin.H{"message": "Notion API Failed!"})
	}

	// TODO slack に送信
	slack.New("YOUR_TOKEN_HERE")

	c.JSON(200, gin.H{
		"message":    "thanksfulness!",
		"thanksList": thanksList, // TODO 動作確認用 あとで消す
	})
}
