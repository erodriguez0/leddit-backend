package api

import (
	db "github.com/erodriguez0/leddit-backend/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server serves HHTP request for out banking service
type Server struct {
	service db.Service
	router  *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(service db.Service) *Server {
	server := &Server{service: service}
	router := gin.Default()
	router.SetTrustedProxies([]string{})

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("user_roles", validUserRole)
	}

	router.POST("/users", server.createUser)
	// router.GET("/users/:id", server.getUser)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
