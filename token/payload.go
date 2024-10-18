package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrTokenInvalid = errors.New("token is invalid")
	ErrTokenExpired = errors.New("token has expired")
)

type Payload struct {
	Id         uuid.UUID
	Username   string
	Created_at time.Time
	Expired_at time.Time
}

// Valid implements jwt.Claims.

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	resp := &Payload{
		Id:         id,
		Username:   username,
		Created_at: time.Now(),
		Expired_at: time.Now().Add(duration),
	}
	return resp, nil
}

func (p *Payload) Valid() error {
	if time.Now().After(p.Expired_at) {
		return errors.New("toke has expired")
	}
	return nil
}
