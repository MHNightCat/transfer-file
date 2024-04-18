package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.StaticFile("/test.mp4", "./static/test.mp4")

	router.Run(":3000")
}
