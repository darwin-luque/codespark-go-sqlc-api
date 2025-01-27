-- name: AddArticleAsFavorite :one
INSERT INTO "favorite" ("user_id", "article_id")
VALUES ($1, $2)
RETURNING *;
