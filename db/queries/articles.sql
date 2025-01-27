-- name: CreateArticle :one
INSERT INTO "article"("title", "body", "description", "slug", "author_id", "status")
  VALUES ($1, $2, $3, $4, sqlc.arg('author_id')::uuid, $5)
RETURNING
  *;

-- name: ListArticles :many
SELECT
  *
FROM
  "article"
WHERE
  "status" = 'published'
  AND (
    "title" ILIKE '%' || coalesce(sqlc.narg('title'), '') || '%' OR
    "slug" ILIKE '%' || coalesce(sqlc.narg('slug'), '') || '%'
  )
LIMIT coalesce(sqlc.narg('limit')::int, 10) OFFSET coalesce(sqlc.narg('offset')::int, 0);
