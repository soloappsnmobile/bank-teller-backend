package routers

import (
	"bank-teller-backend/handlers/admin"
	"bank-teller-backend/handlers/auth"
	"bank-teller-backend/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Enable CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	router.Use(cors.New(config))

	// Set up the routes
	router.POST("/v1/auth/login", auth.Login)
	router.POST("/v1/admin/create-teller", middlewares.TokenAuthMiddleware("Admin"), admin.CreateTeller)
	router.POST("/v1/admin/create-customer", middlewares.TokenAuthMiddleware("Admin", "Teller"), admin.CreateCustomer)
	router.GET("/v1/admin/get-tellers", middlewares.TokenAuthMiddleware("Admin"), admin.GetTellers)

	return router

}
