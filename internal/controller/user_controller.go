package controller

import (
	"net/http"
	"strings"

	"github.com/ElCesarSP/go-api/internal/model"
	"github.com/ElCesarSP/go-api/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var input model.CreateUserInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := c.service.CreateUser(ctx.Request.Context(), input)
	if err != nil {

		if strings.Contains(err.Error(), "Unique constraint failed") {
			ctx.JSON(http.StatusConflict, gin.H{
				"error": "email already exists",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
		"user":    user,
	})
}

func (c *UserController) Login(ctx *gin.Context) {

	var input model.LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := c.service.Login(
		ctx.Request.Context(),
		input,
	)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (c *UserController) Me(ctx *gin.Context) {

	userID := ctx.MustGet("userID").(string)

	user, err := c.service.Me(
		ctx.Request.Context(),
		userID,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch user",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
