package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", fff)
	// fmt.Println("hello there")
	router.Run(":8000")
}

func fff(c *gin.Context) {
}
