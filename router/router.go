package router

import (
	"essential/controller"
	"essential/middleware"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CorsMiddleware(), middleware.RecoveryMiddleware())
	auth := r.Group("v1/auth")
	{
		auth.POST("/register", controller.Register)
		auth.POST("/login", controller.Login)
		auth.GET("/info", middleware.AuthMiddleware(), controller.Info)

	}

	categoryRouter := r.Group("v1/category")
	categoryRouter.Use(middleware.AuthMiddleware())
	categoryController := controller.NewCategoryController()
	{
		categoryRouter.GET("/:id",  categoryController.Show)
		categoryRouter.POST("/post",categoryController.Create)
		categoryRouter.PUT("/:id", categoryController.Update)
		categoryRouter.DELETE("/:id", categoryController.Delete)
	}

	postsRouter := r.Group("/v1/posts")
	postsRouter.Use(middleware.AuthMiddleware())
	postController := controller.NewPostsRouterController()
	{
		postsRouter.GET("/:id",postController.Show)
		postsRouter.POST("/post", postController.Create)
		postsRouter.PUT("/:id", postController.Update)
		postsRouter.DELETE("/:id", postController.Delete)
	}
	return r
}
