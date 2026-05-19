package routes

import (
	"github.com/ElCesarSP/go-api/internal/controller"
	"github.com/ElCesarSP/go-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Initialize(UserController *controller.UserController) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET(
		"/profile",
		middleware.AuthMiddleware(),
		func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "authenticated user",
			})
		},
	)

	router.GET(
		"/me",
		middleware.AuthMiddleware(),
		UserController.Me,
	)

	router.POST("/users", UserController.CreateUser)
	router.POST("/login", UserController.Login)

	return router
}
