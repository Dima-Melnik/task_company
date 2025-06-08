package middleware

import "github.com/gin-gonic/gin"

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.MustGet("role").(string)
		if role != "admin" {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden: Admins only"})
			return
		}
		c.Next()
	}
}
