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

type createPostRequest struct {
	Title         string `json:"title"`
	Url           string `json:"url"`
	Body          string `json:"body"`
	SubledditName string `json:"subleddit_name"`
}

type postSubleddit struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type postResponse struct {
	Title     string        `json:"title"`
	Url       string        `json:"url"`
	Body      string        `json:"body"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	Subleddit postSubleddit `json:"subleddit"`
	User      userResponse  `json:"user"`
}

func newPostResposne(post *db.CreatePostRow) postResponse {
	return postResponse{
		Title:     post.Title,
		Url:       post.Url.String,
		Body:      post.Body.String,
		CreatedAt: post.CreatedAt.Time,
		UpdatedAt: post.UpdatedAt.Time,
		Subleddit: postSubleddit{
			Name:      post.Subleddit.Name,
			CreatedAt: post.Subleddit.CreatedAt.Time,
			UpdatedAt: post.Subleddit.UpdatedAt.Time,
		},
		User: newUserResponse(post.User),
	}
}

func (server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	subleddit, err := server.service.GetSubleddit(ctx, req.SubledditName)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreatePostParams{
		Title:       req.Title,
		Url:         sql.NullString{String: req.Url, Valid: true},
		Body:        sql.NullString{String: req.Body, Valid: true},
		SubledditID: uuid.NullUUID{UUID: subleddit.ID, Valid: true},
		UserID:      uuid.NullUUID{UUID: authPayload.User.Id, Valid: true},
	}

	post, err := server.service.CreatePost(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newPostResposne(&post)

	ctx.JSON(http.StatusOK, rsp)
}
