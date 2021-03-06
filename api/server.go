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
	router *gin.Engine  // 加入router的目的是能在main文件中以特定端口形式启动server
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

	// 下面的接口不用认证
	router.POST("/users", server.createUser)
	router.POST("/users/login",server.loginUser)

	// 生成新的router group，在这个group中的请求都需要认证
	authRoute := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoute.POST("/accounts", server.createAccount)
	authRoute.GET("/accounts/:id", server.getAccount)
	authRoute.GET("/accounts", server.listAccount)
	authRoute.POST("/transfers", server.createTransfer)

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
