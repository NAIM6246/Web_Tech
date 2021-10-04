package controllers

import (
	"encoding/json"
	"net/http"
	"webTech/models/domains"
	"webTech/services"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	IController
}

type UserController struct {
	userService services.IUserService
}

func NewUserController(
	userService services.IUserService) IUserController {
	return &UserController{
		userService: userService,
	}
}

func (h *UserController) Handler(r *gin.Engine) {
	router := r.Group("/users")
	{
		router.GET("/", h.getAll)
		router.POST("/", h.createUser)
	}
}

func (h *UserController) createUser(c *gin.Context) {
	user := domains.User{}
	if parsingErr := json.NewDecoder(c.Request.Body).Decode(&user); parsingErr != nil {
		c.JSON(http.StatusBadRequest, parsingErr)
		return
	}
	createduser, err := h.userService.CreateUser(&user)
}

func (h *UserController) getAll(r *gin.Context) {

}
