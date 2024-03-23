package api

import (
	"net/http"
	"time"

	db "github.com/erodriguez0/leddit-backend/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createSubledditRequest struct {
	Name   string    `json:"name"`
	UserId uuid.UUID `json:"user_id"`
}

type subledditResponse struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newSubledditResponse(subleddit *db.Subleddit) subledditResponse {
	return subledditResponse{
		Name:      subleddit.Name,
		CreatedAt: subleddit.CreatedAt.Time,
		UpdatedAt: subleddit.UpdatedAt.Time,
	}
}

func (server *Server) createSubleddit(ctx *gin.Context) {
	var req createSubledditRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateSubledditParams{
		Name:   req.Name,
		UserID: uuid.NullUUID{UUID: req.UserId, Valid: true},
	}

	// TODO: Check authorization header
	subleddit, err := server.service.CreateSubleddit(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newSubledditResponse(&subleddit)
	ctx.JSON(http.StatusOK, rsp)
}
