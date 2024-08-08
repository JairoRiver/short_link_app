package api

import (
	_ "github.com/JairoRiver/short_link_app/short_link/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (server *Server) setupRouter() {
	router := gin.Default()
	// @title Short Link API
	// @version 1.0
	// @description Testing Swagger APIs.
	// @termsOfService http://swagger.io/terms/// @contact.name API Support
	// @contact.url http://www.swagger.io/support
	// @contact.email support@swagger.io// @securityDefinitions.apiKey JWT
	// @in header
	// @name token// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html// @host localhost:8081
	// @BasePath /v1// @schemes http
	// Swagger documentation
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") //Update the route whe load .env file
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	//Load HTML templates
	router.LoadHTMLGlob("./templates/*")

	//Create Routes
	router.POST("/v1/create", server.handler.CreateLink)

	//Get Routes
	router.GET("/:token", server.handler.GetLink)

	//Availability token check route
	router.POST("/check/:token", server.handler.AvailabilityLink)

	server.router = router
}
