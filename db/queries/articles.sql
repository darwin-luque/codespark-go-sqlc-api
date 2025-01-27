-- name: CreateArticle :one
INSERT INTO "article"("title", "body", "description", "slug", "author_id", "status")
  VALUES ($1, $2, $3, $4, sqlc.arg('author_id')::uuid, $5)
RETURNING
  *;
