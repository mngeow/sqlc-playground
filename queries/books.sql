-- name: GetBook :one
SELECT * FROM books
WHERE id = ? LIMIT 1;

-- name: CreateBook :one
INSERT INTO books (
    name, author_id
) VALUES (
    ?, ?
)
RETURNING *;