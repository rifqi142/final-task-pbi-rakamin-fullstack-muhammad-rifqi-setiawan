package router

import (
	"final-task-pbi-rakamin-fullstack-muhammad-rifqi-setiawan/controller"
	"final-task-pbi-rakamin-fullstack-muhammad-rifqi-setiawan/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.Register)
		userRouter.POST("/login", controller.Login)
		userRouter.PUT("/",middleware.Authentication(), controller.UpdateDataUser)
		userRouter.DELETE("/:UserId",middleware.Authentication(), controller.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("", controller.CreatePhoto)
		photoRouter.GET("/:PhotoId", controller.GetPhotoById)
		photoRouter.PUT("/:PhotoId", controller.UpdatePhoto)
		photoRouter.DELETE("/:PhotoId", controller.DeletePhoto)
	}

	return r
}