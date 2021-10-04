package controllers

import "github.com/gin-gonic/gin"

type IController interface {
	Handler(r *gin.Engine)
}
