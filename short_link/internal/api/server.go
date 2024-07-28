package api

import (
	"github.com/JairoRiver/short_link_app/short_link/internal/api/handler/rest"
	"github.com/gin-gonic/gin"
)

// Server serve a HTTP request.
type Server struct {
	handler *rest.Handler
	router  *gin.Engine
}

// New creates a new Server service HTTP.
func New(handler *rest.Handler) *Server {
	server := Server{handler: handler}
	server.setupRouter()
	return &server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
