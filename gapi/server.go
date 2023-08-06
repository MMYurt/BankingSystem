package gapi

import (
	db "BankingSystem/db/sqlc"
	"BankingSystem/pb"
	"BankingSystem/token"
	"BankingSystem/util"
	"fmt"
)

type Server struct {
	pb.UnimplementedBankingSystemServer
	store      *db.Store
	tokenMaker token.Maker
	config     util.Config
}

// NewServer - Creates a new gRPC server
func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSecretKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{store: store, config: config, tokenMaker: tokenMaker}

	return server, nil
}
