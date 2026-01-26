package routes

import (
	"example/Go-Backend/controllers"

	"github.com/gin-gonic/gin"
)

// MUST start with Capital R
func RegisterProductRoutes(router *gin.Engine) {
	product := router.Group("/api/products")
	{
		product.GET("", controllers.GetAllProducts)
		product.GET("/:id", controllers.GetProductByID)
		product.POST("", controllers.CreateProduct)
		product.PUT("/:id", controllers.UpdateProduct)
		product.DELETE("/:id", controllers.DeleteProduct)
	}
}
