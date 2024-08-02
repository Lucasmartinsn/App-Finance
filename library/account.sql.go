// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: account.sql

package library

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAccounts = `-- name: CreateAccounts :one
INSERT INTO accounts (
    user_id, categories_id, title, type, description, value, date
) VALUES (
    $1,$2,$3,$4,$5,$6,$7
) RETURNING id, user_id, categories_id, title, type, description, value, date, create_at
`

type CreateAccountsParams struct {
	UserID       pgtype.UUID
	CategoriesID pgtype.UUID
	Title        string
	Type         string
	Description  string
	Value        int32
	Date         pgtype.Date
}

func (q *Queries) CreateAccounts(ctx context.Context, arg CreateAccountsParams) (Account, error) {
	row := q.db.QueryRow(ctx, createAccounts,
		arg.UserID,
		arg.CategoriesID,
		arg.Title,
		arg.Type,
		arg.Description,
		arg.Value,
		arg.Date,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CategoriesID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.Value,
		&i.Date,
		&i.CreateAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteAccount, id)
	return err
}

const getAccountGraph = `-- name: GetAccountGraph :one
SELECT COUNT(*) as sum_value FROM accounts
WHERE user_id = $1 and type = $2
`

type GetAccountGraphParams struct {
	UserID pgtype.UUID
	Type   string
}

func (q *Queries) GetAccountGraph(ctx context.Context, arg GetAccountGraphParams) (int64, error) {
	row := q.db.QueryRow(ctx, getAccountGraph, arg.UserID, arg.Type)
	var sum_value int64
	err := row.Scan(&sum_value)
	return sum_value, err
}

const getAccountReports = `-- name: GetAccountReports :one
SELECT SUM(value) as sum_value FROM accounts
WHERE user_id = $1 and type = $2
`

type GetAccountReportsParams struct {
	UserID pgtype.UUID
	Type   string
}

func (q *Queries) GetAccountReports(ctx context.Context, arg GetAccountReportsParams) (int64, error) {
	row := q.db.QueryRow(ctx, getAccountReports, arg.UserID, arg.Type)
	var sum_value int64
	err := row.Scan(&sum_value)
	return sum_value, err
}

const getAccounts = `-- name: GetAccounts :many
SELECT id, user_id, categories_id, title, type, description, value, date, create_at FROM accounts WHERE user_id = $1 AND type = $2 AND title LIKE $3 AND description LIKE $4 AND date LIKE $5
`

type GetAccountsParams struct {
	UserID      pgtype.UUID
	Type        string
	Title       string
	Description string
	Date        pgtype.Date
}

func (q *Queries) GetAccounts(ctx context.Context, arg GetAccountsParams) ([]Account, error) {
	rows, err := q.db.Query(ctx, getAccounts,
		arg.UserID,
		arg.Type,
		arg.Title,
		arg.Description,
		arg.Date,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CategoriesID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.Value,
			&i.Date,
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

const getAccountsById = `-- name: GetAccountsById :one
SELECT id, user_id, categories_id, title, type, description, value, date, create_at FROM accounts WHERE id = $1
`

func (q *Queries) GetAccountsById(ctx context.Context, id pgtype.UUID) (Account, error) {
	row := q.db.QueryRow(ctx, getAccountsById, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CategoriesID,
		&i.Title,
		&i.Type,
		&i.Description,
		&i.Value,
		&i.Date,
		&i.CreateAt,
	)
	return i, err
}

const getAccountsFull = `-- name: GetAccountsFull :many
SELECT 
    a.id,
    a.user_id,
    a.title,
    a.type,
    a.description,
    a.value,
    a.date,
    a.create_at,
    c.title as categories_title
FROM accounts a LEFT JOIN categories c ON c.id = a.categories_id
WHERE a.user_id = $1 AND a.type = $2 AND a.title LIKE $3 
AND a.categories_id = $4 AND a.description LIKE $5 AND a.date = $6
`

type GetAccountsFullParams struct {
	UserID       pgtype.UUID
	Type         string
	Title        string
	CategoriesID pgtype.UUID
	Description  string
	Date         pgtype.Date
}

type GetAccountsFullRow struct {
	ID              pgtype.UUID
	UserID          pgtype.UUID
	Title           string
	Type            string
	Description     string
	Value           int32
	Date            pgtype.Date
	CreateAt        pgtype.Timestamptz
	CategoriesTitle pgtype.Text
}

func (q *Queries) GetAccountsFull(ctx context.Context, arg GetAccountsFullParams) ([]GetAccountsFullRow, error) {
	rows, err := q.db.Query(ctx, getAccountsFull,
		arg.UserID,
		arg.Type,
		arg.Title,
		arg.CategoriesID,
		arg.Description,
		arg.Date,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAccountsFullRow
	for rows.Next() {
		var i GetAccountsFullRow
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Title,
			&i.Type,
			&i.Description,
			&i.Value,
			&i.Date,
			&i.CreateAt,
			&i.CategoriesTitle,
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

const updateAccounts = `-- name: UpdateAccounts :exec
UPDATE accounts SET 
    title = $2,
    description = $3,
    value = $4
WHERE id = $1
`

type UpdateAccountsParams struct {
	ID          pgtype.UUID
	Title       string
	Description string
	Value       int32
}

func (q *Queries) UpdateAccounts(ctx context.Context, arg UpdateAccountsParams) error {
	_, err := q.db.Exec(ctx, updateAccounts,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Value,
	)
	return err
}
