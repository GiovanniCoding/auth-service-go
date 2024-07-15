-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (email, password_hash)
VALUES ($1, $2)
RETURNING *;

-- name: UserEmailExist :one
SELECT EXISTS (
    SELECT 1
    FROM users
    WHERE email = $1
) AS exists;
