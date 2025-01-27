-- name: AddComment :one
INSERT INTO "comment"("article_id", "author_id", "body")
  VALUES ($1, $2, $3)
RETURNING
  *;

-- name: ListCommentsForArticleBySlug :many
SELECT
  "comment"."id",
  "comment"."body",
  "comment"."created_at",
  "comment"."updated_at",
  "comment"."article_id",
  "comment"."author_id"
FROM
  "comment"
INNER JOIN "article" ON "comment"."article_id" = "article"."id"
WHERE
  "article"."slug" = $1;

-- name: DeleteComment :exec
DELETE FROM "comment"
WHERE
  "id" = $1::uuid;

-- name: GetComment :one
SELECT
  *
FROM
  "comment"
WHERE
  "id" = @comment_id::uuid;
