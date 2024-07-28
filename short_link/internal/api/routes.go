package api

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) setupRouter() {
	router := gin.Default()

	//Create Routes
	router.POST("/v1/create", server.handler.CreateLink)

	server.router = router
}
