package controllers

import "github.com/gin-gonic/gin"

// type IUserController interface{}

type UserController struct{}

func (h *UserController) Handler(r *gin.Engine) {
	router := r.Group("/users")
	{
		router.GET("/")
		router.POST("/")
	}
}
