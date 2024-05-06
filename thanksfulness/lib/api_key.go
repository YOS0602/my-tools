package lib

import (
	"errors"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func ApiKeyMiddleware(c *gin.Context) {
	thanksfulnessApiKey, ok := os.LookupEnv("THANKSFULNESS_API_KEY")
	if !ok {
		log.Println("Set ENV 'THANKSFULNESS_API_KEY'")
		c.AbortWithStatusJSON(500, gin.H{"message": "internal server error"})
	}

	valid, err := CheckApiKey(c.GetHeader("x-api-key"), thanksfulnessApiKey)
	if !valid || err != nil {
		c.AbortWithStatusJSON(403, gin.H{"message": "invalid api key"})
	}
}

func CheckApiKey(reqApiKey string, thanksfulnessApiKey string) (bool, error) {
	if reqApiKey != thanksfulnessApiKey {
		err := errors.New("invalid api key")
		return false, err
	}
	return true, nil
}
