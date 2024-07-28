package api

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) setupRouter() {
	router := gin.Default()

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
