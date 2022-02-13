package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/oaraujocesar/simple_bank/db/sqlc"
)

// Serves HTTP requests for our banking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Creates a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	router.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router

	return server
}

// Runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
