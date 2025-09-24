package routes

import (
	"spad_be/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", controllers.CreateUser)
		userRoutes.GET("", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUserByID)
	}
}
