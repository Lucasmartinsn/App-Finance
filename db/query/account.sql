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

-- name: GetAccountsFull :many
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
AND a.categories_id = $4 AND a.description LIKE $5 AND a.date = $6;

-- name: UpdateAccounts :exec
UPDATE accounts SET 
    title = $2,
    description = $3,
    value = $4
WHERE id = $1;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;

-- name: GetAccountReports :one
SELECT SUM(value) as sum_value FROM accounts
WHERE user_id = $1 and type = $2;

-- name: GetAccountGraph :one
SELECT COUNT(*) as sum_value FROM accounts
WHERE user_id = $1 and type = $2;