package api

import (
	db "BankingSystem/db/sqlc"
	"BankingSystem/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store *db.Store) *Server {
	config := util.Config{
		TokenSecretKey:      "12345678123456781234567812345678",
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}
