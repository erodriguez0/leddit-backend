package api

import (
	"fmt"

	db "github.com/erodriguez0/leddit-backend/db/sqlc"
	"github.com/erodriguez0/leddit-backend/token"
	"github.com/erodriguez0/leddit-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HHTP request for out banking service
type Server struct {
	service    db.Service
	router     *gin.Engine
	tokenMaker token.Maker
	config     util.Config
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, service db.Service) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	server := &Server{
		config:     config,
		service:    service,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()
	router.SetTrustedProxies([]string{})
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("user_roles", validUserRole)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/register", server.registerUser)
			auth.POST("/login", server.loginUser)
		}

		users := v1.Group("/users")
		{
			users.GET("/:username", server.getUser)
		}

		subleddits := v1.Group("/subleddits")
		{
			subleddits.POST("/", authMiddleware(server.tokenMaker), server.createSubleddit)
			subleddits.GET("/:name", server.getSubleddit)
		}

		posts := v1.Group("/posts")
		{
			posts.POST("/", authMiddleware(server.tokenMaker), server.createPost)
		}
	}

	server.router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
