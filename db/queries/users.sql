-- name: CreateUser :one
INSERT INTO "user" ("username", "email", "bio", "image", "password_hash")
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM "user"
WHERE "email" = $1 LIMIT 1;
