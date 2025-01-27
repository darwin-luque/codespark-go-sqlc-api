-- name: FollowUser :one
INSERT INTO "follow"("following_user_id", "followed_user_id")
  VALUES (sqlc.arg('following_user_id')::uuid, sqlc.arg('followed_user_id')::uuid)
RETURNING
  *;

-- name: UnfollowUser :exec
DELETE FROM "follow"
WHERE "following_user_id" = sqlc.arg('following_user_id')::uuid
  AND "followed_user_id" = sqlc.arg('followed_user_id')::uuid;
