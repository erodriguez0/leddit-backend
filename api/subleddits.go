package api

import (
	"net/http"
	"time"

	db "github.com/erodriguez0/leddit-backend/db/sqlc"
	"github.com/erodriguez0/leddit-backend/token"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createSubledditRequest struct {
	Name string `json:"name"`
}

type subledditResponse struct {
	Name      string       `json:"name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	User      userResponse `json:"user"`
}

func newSubledditResponse(subleddit *db.CreateSubledditRow) subledditResponse {
	return subledditResponse{
		Name:      subleddit.Name,
		CreatedAt: subleddit.CreatedAt.Time,
		UpdatedAt: subleddit.UpdatedAt.Time,
		User:      newUserResponse(subleddit.User),
	}
}

func (server *Server) createSubleddit(ctx *gin.Context) {
	var req createSubledditRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateSubledditParams{
		Name:   req.Name,
		UserID: uuid.NullUUID{UUID: authPayload.User.Id, Valid: false},
	}

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
