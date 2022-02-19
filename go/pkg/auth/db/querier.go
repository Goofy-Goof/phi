// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"

	"github.com/gofrs/uuid"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByName(ctx context.Context, name string) (User, error)
}

var _ Querier = (*Queries)(nil)
