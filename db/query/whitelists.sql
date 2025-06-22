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
    AND ((sqlc.narg ('platform_user_id')::string IS NULL
            OR platform_user_id = sqlc.narg ('platform_user_id'))
        OR (sqlc.narg ('platform_user_name')::string IS NULL
            OR platform_user_name = sqlc.narg ('platform_user_name'))
        OR (sqlc.narg ('platform_user_login')::string IS NULL
            OR platform_user_login = sqlc.narg ('platform_user_login'))
        OR (sqlc.narg ('user_id')::uuid IS NULL
            OR user_id = sqlc.narg ('user_id')));

-- name: WhitelistDelete :execrows
DELETE FROM public.whitelist
WHERE platform = $1
    AND ((sqlc.narg ('platform_user_id')::string IS NULL
            OR platform_user_id = sqlc.narg ('platform_user_id'))
        AND (sqlc.narg ('platform_user_name')::string IS NULL
            OR platform_user_name = sqlc.narg ('platform_user_name'))
        AND (sqlc.narg ('platform_user_login')::string IS NULL
            OR platform_user_login = sqlc.narg ('platform_user_login'))
        AND (sqlc.narg ('user_id')::uuid IS NULL
            OR user_id = sqlc.narg ('user_id')));

