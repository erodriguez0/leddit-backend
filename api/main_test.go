package api

import (
	"os"
	"testing"
	"time"

	db "github.com/erodriguez0/leddit-backend/db/sqlc"
	"github.com/erodriguez0/leddit-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, service db.Service) *Server {
	config := util.Config{
		TokenSymmetricKey:  util.RandomString(32),
		AccessTokenDuraion: time.Minute,
	}

	server, err := NewServer(config, service)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
