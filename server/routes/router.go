package routes

import (
	"github.com/douglasandradeee/web_api_gin/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/:id", controllers.FindBookByID)
			books.GET("/", controllers.FindAllBooks)
			books.POST("/post", controllers.CreateBook)
			books.PUT("/update", controllers.UpdadeBook)
			books.DELETE("/:id", controllers.DeleteBookByID)
		}
	}
	return router
}
