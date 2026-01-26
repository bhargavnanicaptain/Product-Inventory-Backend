package main

import (
	"example/Go-Backend/config"
	"example/Go-Backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Connect to MySQL
	config.ConnectDB()

	// 2. Create Gin router
	router := gin.Default()

	// 3. CORS (Required for React)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// 4. Register routes
	routes.RegisterProductRoutes(router)

	// 5. Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "UP",
			"service": "product-backend",
		})
	})

	// 6. Start server
	router.Run(":8080")
}
