package api

import (
	"fmt"
	db "github.com/Ruadgedy/simplebank/db/sqlc"
	"github.com/Ruadgedy/simplebank/token"
	"github.com/Ruadgedy/simplebank/util"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)
import "github.com/gin-gonic/gin"

// Server serves HTTP requests for our banking service
type Server struct {
	config util.Config
	store  db.Store
	tokenMaker token.Maker
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config,store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil,fmt.Errorf("cannot create token maker:%v",err)
	}
	server := &Server{
		config: config,
		store: store,
		tokenMaker:tokenMaker,
	}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// add routes to router
	router.POST("/users", server.createUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.POST("/transfers", server.createTransfer)

	server.router = router
	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Start starts the Http server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
