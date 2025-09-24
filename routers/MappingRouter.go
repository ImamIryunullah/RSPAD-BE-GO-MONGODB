package routes

import (
	"spad_be/controllers"

	"github.com/gin-gonic/gin"
)

func MappingRouter(r *gin.RouterGroup) {
	mapRouter := r.Group("/mapping")
	{
		mapRouter.POST("", controllers.CreateMapping)
		mapRouter.GET("", controllers.GetMappings)
		mapRouter.PUT("/:id", controllers.UpdateMapping)
		mapRouter.DELETE("/:id", controllers.DeleteMapping)
	}
}
