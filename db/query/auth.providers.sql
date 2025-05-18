-- name: AuthProviderGetById :one
SELECT *
FROM auth.providers
WHERE id = $1 AND provider = $2;

-- name: AuthProviderGetByUserId :one
SELECT *
FROM auth.providers
WHERE user_id = $1 AND provider = $2;

-- name: AuthProviderGetByProviderUserId :one
SELECT *
FROM auth.providers
WHERE provider_user_id = $1 AND provider = $2;

-- name: CreateAuthProvider :one
INSERT INTO auth.providers (
  user_id,
  provider_user_id,
  provider,
  access_token,
  refresh_token,
  access_type
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
) RETURNING id;

-- name: AuthProviderUpdateTokens :execrows
UPDATE auth.providers
SET
access_token = COALESCE($1, access_token),
-- refresh_token = COALESCE($2, refresh_token),
refresh_token = sqlc.narg('refresh_token'),
updated_at = CURRENT_TIMESTAMP
WHERE id = $2;

