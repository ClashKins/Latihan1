package router

import (
	"LATIHAN1/controllers"
	"LATIHAN1/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}
	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetAllPhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.CommentAuthorization())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetAllComments)
		commentRouter.GET("/:commentId", controllers.GetOneComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}
	socialMediaRouter := r.Group("/social-medias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", controllers.GetAllSocialMedia)
		socialMediaRouter.GET("/:commentId", controllers.GetOneSocialMedia)
		socialMediaRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteSocialMedia)
	}
	return r
}