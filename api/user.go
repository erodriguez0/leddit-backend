package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required,min=8"`
	Avatar   string `json:"avatar" binding:"url"`
	UserRole string `json:"user_role" binding:"required,user_roles"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, req)
}
