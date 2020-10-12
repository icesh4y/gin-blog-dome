package router

import (
	"essential/controller"
	"essential/middleware"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("v1/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
		auth.GET("/info", middleware.AuthMiddleware(), controller.Info)

	}

	category := r.Group("v1/category")
	categorycontroller := controller.NewCategoryController()
	{
		category.GET("/:id", categorycontroller.Show)
		category.POST("/post", categorycontroller.Create)
		category.PUT("/:id", categorycontroller.Update)
		category.DELETE("/:id", categorycontroller.Delete)
	}

	return r
}
