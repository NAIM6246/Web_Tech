package main

import (
	"fmt"
	"webTech/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", fff)
	fmt.Println("hello there")
	router.Run(":8000")
}

func fff(c *gin.Context) {
	e := services.UserService{}
	fmt.Println(e.Hello())
	fmt.Println("hi from chi")
}
