// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package db

import (
	"context"

	"github.com/gofrs/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  id,
  tink_id
) VALUES (
  $1,
  $2
) ON CONFLICT DO NOTHING RETURNING id
`

type CreateUserParams struct {
	ID     uuid.UUID `json:"id"`
	TinkID string    `json:"tink_id"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (uuid.UUID, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser, arg.ID, arg.TinkID)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT
  id, tink_id
FROM
  users
WHERE
  id = $1::uuid
`

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.queryRow(ctx, q.getUserByIDStmt, getUserByID, id)
	var i User
	err := row.Scan(&i.ID, &i.TinkID)
	return i, err
}