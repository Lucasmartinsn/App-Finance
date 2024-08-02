-- name: CreateUser :one
INSERT INTO users (
    username, password, email
) VALUES (
    $1,$2,$3
) RETURNING *;

-- name: GetUser :many
SELECT * FROM users WHERE username LIKE $1;

-- name: GetUserById :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users SET
    username = $2,
    password = $3,
    email = $4
WHERE id = $1;

-- name: UpdateUserPass :exec
UPDATE users SET
    password = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;