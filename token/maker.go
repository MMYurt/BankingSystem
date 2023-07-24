package token

import "time"

// Maker - interface for managing tokens
type Maker interface {
	// CreateToken - creates a new token
	CreateToken(username string, duration time.Duration) (string, error)

	// VerifyToken - checks if the token is valid
	VerifyToken(token string) (*Payload, error)
}
