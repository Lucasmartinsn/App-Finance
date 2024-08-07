// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: category.sql

package library

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories (
    user_id, title, type, description
) VALUES (
    $1, $2, $3, $4
) RETURNING id, user_id, title, type, description, create_at
`

type CreateCategoryParams struct {
	UserID      pgtype.UUID
	Title       string
	Type        string
	Description string
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRow(ctx, createCategory,
		arg.UserID,
		arg.Title,
		arg.Type,
		arg.Description,
	)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.CreateAt,
	)
	return i, err
}

const getCategory = `-- name: GetCategory :many
SELECT id, user_id, title, type, description, create_at FROM categories WHERE user_id = $1 AND type = $2 AND title LIKE $3 AND description LIKE $4
`

type GetCategoryParams struct {
	UserID      pgtype.UUID
	Type        string
	Title       string
	Description string
}

func (q *Queries) GetCategory(ctx context.Context, arg GetCategoryParams) ([]Category, error) {
	rows, err := q.db.Query(ctx, getCategory,
		arg.UserID,
		arg.Type,
		arg.Title,
		arg.Description,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.CreateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCategoryById = `-- name: GetCategoryById :one
SELECT id, user_id, title, type, description, create_at FROM categories WHERE id = $1
`

func (q *Queries) GetCategoryById(ctx context.Context, id pgtype.UUID) (Category, error) {
	row := q.db.QueryRow(ctx, getCategoryById, id)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.CreateAt,
	)
	return i, err
}

const getCategoryByUserId = `-- name: GetCategoryByUserId :many
SELECT id, user_id, title, type, description, create_at FROM categories WHERE user_id = $1 AND type LIKE $2
`

type GetCategoryByUserIdParams struct {
	UserID pgtype.UUID
	Type   string
}

func (q *Queries) GetCategoryByUserId(ctx context.Context, arg GetCategoryByUserIdParams) ([]Category, error) {
	rows, err := q.db.Query(ctx, getCategoryByUserId, arg.UserID, arg.Type)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.CreateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategory = `-- name: UpdateCategory :exec
UPDATE categories SET 
    title = $2,
    type = $3,
    description = $4
WHERE id = $1
`

type UpdateCategoryParams struct {
	ID          pgtype.UUID
	Title       string
	Type        string
	Description string
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.Exec(ctx, updateCategory,
		arg.ID,
		arg.Title,
		arg.Type,
		arg.Description,
	)
	return err
}
