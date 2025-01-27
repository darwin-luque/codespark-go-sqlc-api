-- name: CreateUser :one
INSERT INTO "user"("username", "email", "bio", "image", "password_hash")
  VALUES ($1, $2, $3, $4, $5)
RETURNING
  *;

-- name: GetUserByEmail :one
SELECT
  "user"."id",
  "user"."username",
  "user"."email",
  "user"."bio",
  "user"."image",
  "user"."password_hash",
  "user"."created_at",
  "user"."updated_at",
  COUNT(DISTINCT "followers"."following_user_id") AS "followers_count",
  COUNT(DISTINCT "following"."followed_user_id") AS "following_count"
FROM
  "user"
  LEFT JOIN "follow" AS "followers" ON "followers"."followed_user_id" = "user"."id"
  LEFT JOIN "follow" AS "following" ON "following"."following_user_id" = "user"."id"
WHERE
  "email" = $1
GROUP BY "user"."id"
LIMIT 1;

-- name: GetUserByUsername :one
SELECT
  "user"."id",
  "user"."username",
  "user"."email",
  "user"."bio",
  "user"."image",
  "user"."password_hash",
  "user"."created_at",
  "user"."updated_at",
  COUNT(DISTINCT "followers"."following_user_id") AS "followers_count",
  COUNT(DISTINCT "following"."followed_user_id") AS "following_count"
FROM
  "user"
  LEFT JOIN "follow" AS "followers" ON "followers"."followed_user_id" = "user"."id"
  LEFT JOIN "follow" AS "following" ON "following"."following_user_id" = "user"."id"
WHERE
  "username" = $1
GROUP BY "user"."id"
LIMIT 1;

