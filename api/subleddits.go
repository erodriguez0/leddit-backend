package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/erodriguez0/leddit-backend/db/sqlc"
	"github.com/erodriguez0/leddit-backend/token"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createSubledditRequest struct {
	Name string `uri:"name" binding:"required,alphanum"`
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

	// ctx.JSON(http.StatusOK, authPayload.User.Id)
	// return

	arg := db.CreateSubledditParams{
		Name:   req.Name,
		UserID: uuid.NullUUID{UUID: authPayload.User.Id, Valid: true},
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

type getSubledditRequest struct {
	Name string `uri:"name" binding:"required,alphanum"`
}

// TODO: Extend to include posts
type getSubledditResponse struct {
	Name      string       `json:"name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	User      userResponse `json:"user"`
}

func newGetSubledditResponse(subleddit *db.GetSubledditRow) getSubledditResponse {
	return getSubledditResponse{
		Name:      subleddit.Name,
		CreatedAt: subleddit.CreatedAt.Time,
		UpdatedAt: subleddit.UpdatedAt.Time,
		User:      newUserResponse(subleddit.User),
	}
}

func (server *Server) getSubleddit(ctx *gin.Context) {
	var req getSubledditRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	subleddit, err := server.service.GetSubleddit(ctx, req.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newGetSubledditResponse(&subleddit)

	ctx.JSON(http.StatusOK, rsp)
}
