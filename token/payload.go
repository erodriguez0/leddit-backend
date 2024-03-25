package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type tokenUser struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

// Payload contains the payload data of the token
type Payload struct {
	Id        uuid.UUID `json:"jti"`
	IssuedAt  int64     `json:"iat"`
	ExpiresAt int64     `json:"exp"`
	User      tokenUser `json:"user"`
}

func NewPayload(user_id uuid.UUID, username string, duration time.Duration) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		Id:        tokenId,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(duration).Unix(),
		User: tokenUser{
			Id:       user_id,
			Username: username,
		},
	}

	return payload, nil
}

// Valid checks if the token payload is valid
func (payload *Payload) Valid() error {
	if time.Now().After(time.Unix(payload.ExpiresAt, 0)) {
		return ErrExpiredToken
	}

	return nil
}
