-- name: WhitelistCreate :one
INSERT INTO public.whitelist (platform, platform_user_id, platform_user_name, platform_user_login, user_id)
    VALUES ($1, $2, $3, $4, $5)
RETURNING
    *;

-- name: WhitelistGetOne :one
SELECT
    *
FROM
    public.whitelist
WHERE
    platform = $1
    AND ((sqlc.narg ('platform_user_id')::varchar(100) IS NULL
            OR platform_user_id = sqlc.narg ('platform_user_id'))
        OR (sqlc.narg ('platform_user_name')::varchar(100) IS NULL
            OR platform_user_name = sqlc.narg ('platform_user_name'))
        OR (sqlc.narg ('platform_user_login')::varchar(100) IS NULL
            OR platform_user_login = sqlc.narg ('platform_user_login'))
        OR (sqlc.narg ('user_id')::uuid IS NULL
            OR user_id = sqlc.narg ('user_id')));

-- name: WhitelistUpdate :one
UPDATE
    public.whitelist
SET
    platform = coalesce(sqlc.narg('platform'), platform),
    platform_user_id = coalesce($2, platform_user_id),
    platform_user_name = coalesce($3, platform_user_name),
    platform_user_login = coalesce($4, platform_user_login),
    user_id = coalesce($5, user_id)
WHERE
    id = $1
RETURNING
    *;

-- name: WhitelistDelete :execrows
DELETE FROM public.whitelist
WHERE platform = $1
    AND ((sqlc.narg ('platform_user_id')::varchar(100) IS NULL
            OR platform_user_id = sqlc.narg ('platform_user_id'))
        AND (sqlc.narg ('platform_user_name')::varchar(100) IS NULL
            OR platform_user_name = sqlc.narg ('platform_user_name'))
        AND (sqlc.narg ('platform_user_login')::varchar(100) IS NULL
            OR platform_user_login = sqlc.narg ('platform_user_login'))
        AND (sqlc.narg ('user_id')::uuid IS NULL
            OR user_id = sqlc.narg ('user_id')));

