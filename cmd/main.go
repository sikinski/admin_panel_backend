package main

import (
	"adminka/api"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Hash implements root.Hash
type Hash struct{}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		} else {
			c.Next()
		}
	}
}

// ___________________________________ IT Tasks ___________________________________

// ___________________________________ / IT Tasks ___________________________________

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.POST("/auth", api.Auth)
	router.POST("/createTaskIT", api.CreateTaskIT)

	router.Run("localhost:8080")

}
