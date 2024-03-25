// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: subleddits.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createSubleddit = `-- name: CreateSubleddit :one
WITH Subleddit AS (
    INSERT INTO "subleddits" (
        name,
        user_id
    ) VALUES (
        $1, $2
    ) RETURNING id, name, user_id, created_at, updated_at
) SELECT subleddit.id, subleddit.name, subleddit.user_id, subleddit.created_at, subleddit.updated_at, users.id, users.username, users.password, users.avatar, users.role, users.created_at, users.updated_at FROM Subleddit JOIN users ON Subleddit.user_id = users.id
`

type CreateSubledditParams struct {
	Name   string        `json:"name"`
	UserID uuid.NullUUID `json:"user_id"`
}

type CreateSubledditRow struct {
	ID        uuid.UUID     `json:"id"`
	Name      string        `json:"name"`
	UserID    uuid.NullUUID `json:"user_id"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
	User      User          `json:"user"`
}

func (q *Queries) CreateSubleddit(ctx context.Context, arg CreateSubledditParams) (CreateSubledditRow, error) {
	row := q.db.QueryRowContext(ctx, createSubleddit, arg.Name, arg.UserID)
	var i CreateSubledditRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.User.ID,
		&i.User.Username,
		&i.User.Password,
		&i.User.Avatar,
		&i.User.Role,
		&i.User.CreatedAt,
		&i.User.UpdatedAt,
	)
	return i, err
}

const deleteSubleddit = `-- name: DeleteSubleddit :exec
DELETE FROM "subleddits"
WHERE name = $1
`

func (q *Queries) DeleteSubleddit(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, deleteSubleddit, name)
	return err
}

const getSubleddit = `-- name: GetSubleddit :one
SELECT subleddits.id, subleddits.name, subleddits.user_id, subleddits.created_at, subleddits.updated_at, users.id, users.username, users.password, users.avatar, users.role, users.created_at, users.updated_at
FROM subleddits
JOIN users ON users.id = subleddits.user_id
WHERE subleddits.name = $1
`

type GetSubledditRow struct {
	ID        uuid.UUID     `json:"id"`
	Name      string        `json:"name"`
	UserID    uuid.NullUUID `json:"user_id"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
	User      User          `json:"user"`
}

func (q *Queries) GetSubleddit(ctx context.Context, name string) (GetSubledditRow, error) {
	row := q.db.QueryRowContext(ctx, getSubleddit, name)
	var i GetSubledditRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.User.ID,
		&i.User.Username,
		&i.User.Password,
		&i.User.Avatar,
		&i.User.Role,
		&i.User.CreatedAt,
		&i.User.UpdatedAt,
	)
	return i, err
}
