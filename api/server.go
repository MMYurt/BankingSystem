package api

import (
	db "BankingSystem/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Server for banking services
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer - Creates a new HTTP server
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start - Starts the server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
