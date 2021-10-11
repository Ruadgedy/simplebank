package api

import (
	db "github.com/Ruadgedy/simplebank/db/sqlc"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)
import "github.com/gin-gonic/gin"

// Server serves HTTP requests for our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	if v,ok := binding.Validator.Engine().(*validator.Validate); ok{
		v.RegisterValidation("currency",validCurrency)
	}

	// add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id",server.getAccount)
	router.GET("/accounts",server.listAccount)
	router.POST("/transfers",server.createTransfer)

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Start starts the Http server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}


