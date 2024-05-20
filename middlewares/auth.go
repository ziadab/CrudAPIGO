package middlewares

import "github.com/gin-gonic/gin"

func AuthRequired(c *gin.Context) {

	token := c.GetHeader("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(401, gin.H{"error": "API token required"})
		return
	}

	if token != "heey" {
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
		return
	}

	c.Next()

}
