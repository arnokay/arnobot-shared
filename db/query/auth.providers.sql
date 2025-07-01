-- name: AuthProviderGet :one
SELECT *
FROM auth.providers
WHERE 
(sqlc.narg('user_id')::uuid IS NULL OR user_id = sqlc.narg('user_id')::uuid) AND
(sqlc.narg('provider_user_id')::varchar(100) IS NULL OR provider_user_id = sqlc.narg('provider_user_id')::varchar(100)) AND
provider = $1;

-- name: AuthProviderGetByUserId :one
SELECT *
FROM auth.providers
WHERE user_id = $1 AND provider = $2;

-- name: AuthProviderGetByProviderUserId :one
SELECT *
FROM auth.providers
WHERE provider_user_id = $1 AND provider = $2;

-- name: AuthProviderCreate :one
INSERT INTO auth.providers (
  user_id,
  provider_user_id,
  provider,
  access_token,
  refresh_token,
  access_type,
  scopes
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7
) RETURNING id;

-- name: AuthProviderUpdateTokens :execrows
UPDATE auth.providers
SET
access_token = $1,
refresh_token = $2,
updated_at = CURRENT_TIMESTAMP
WHERE id = $3;

