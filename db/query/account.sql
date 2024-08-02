-- name: CreateAccounts :one
INSERT INTO accounts (
    user_id, categories_id, title, type, description, value, date
) VALUES (
    $1,$2,$3,$4,$5,$6,$7
) RETURNING *;

-- name: GetAccountsById :one
SELECT * FROM accounts WHERE id = $1;

-- name: GetAccounts :many
SELECT * FROM accounts WHERE user_id = $1 AND type = $2 AND title LIKE $3 AND description LIKE $4 AND date LIKE $5;

-- name: UpdateAccounts :exec
UPDATE accounts SET 
    title = $2,
    type = $3,
    description = $4,
    value = $5,
    date = $6
WHERE id = $1;
