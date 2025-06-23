CREATE SCHEMA IF NOT EXISTS public;

CREATE TYPE public.user_status AS ENUM (
    'active',
    'banned',
    'deactivated',
    'deleted'
);

CREATE TYPE public.platform AS ENUM (
    'twitch'
);

CREATE TABLE public.users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    username varchar(50) NOT NULL DEFAULT '',
    status public.user_status NOT NULL DEFAULT 'active',
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE public.user_platform_accounts (
    platform public.platform NOT NULL,
    platform_user_id varchar(100) NOT NULL,
    platform_user_name varchar(100) NOT NULL,
    platform_user_login varchar(100) NOT NULL,
    user_id uuid NOT NULL,
    PRIMARY KEY (platform, platform_user_id),
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE public.whitelist (
    platform public.platform NOT NULL,
    platform_user_id varchar(100),
    platform_user_name varchar(100),
    platform_user_login varchar(100),
    user_id uuid,
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE public.blacklist (
    platform public.platform NOT NULL,
    platform_user_id varchar(100),
    platform_user_name varchar(100),
    platform_user_login varchar(100),
    user_id uuid,
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT
);

