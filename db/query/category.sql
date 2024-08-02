-- name: CreateCategory :one
INSERT INTO categories (
    user_id, title, type, description
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetCategoryById :one
SELECT * FROM categories WHERE id = $1;

-- name: GetCategory :many
SELECT * FROM categories WHERE user_id = $1 AND type = $2 AND title LIKE $3 AND description LIKE $4;

-- name: UpdateCategory :exec
UPDATE categories SET 
    title = $2,
    type = $3,
    description = $4
WHERE id = $1;
