-- name: AddArticleAsFavorite :one
INSERT INTO "favorite" ("user_id", "article_id")
VALUES ($1, $2)
RETURNING *;

-- name: RemoveArticleAsFavorite :exec
DELETE FROM "favorite"
WHERE "user_id" = $1 AND "article_id" = $2;
