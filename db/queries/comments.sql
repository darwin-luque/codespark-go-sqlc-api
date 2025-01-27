-- name: AddComment :one
INSERT INTO "comment"("article_id", "author_id", "body")
  VALUES ($1, $2, $3)
RETURNING
  *;
