package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Tạo router Gin
	router := gin.Default()

	// Định nghĩa route GET gốc "/"
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Gin!",
		})
	})
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello,helloo!",
		})
	})
	// Chạy server tại cổng 8080
	router.Run(":4000")
}
