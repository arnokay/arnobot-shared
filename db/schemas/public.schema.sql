CREATE SCHEMA IF NOT EXISTS public;

CREATE TYPE public.user_status AS ENUM (
    'active',
    'banned',
    'deactivated',
    'deleted'
);

CREATE TABLE public.supported_platforms (
    platform varchar(50) PRIMARY KEY
);

CREATE TABLE public.users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    username varchar(50) NOT NULL DEFAULT '',
    status public.user_status NOT NULL DEFAULT 'active',
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE public.user_platform_accounts (
    platform varchar(50) NOT NULL,
    platform_user_id varchar(100) NOT NULL,
    platform_user_name varchar(100) NOT NULL,
    platform_user_login varchar(100) NOT NULL,
    user_id uuid NOT NULL,
    PRIMARY KEY (platform, platform_user_id),
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT,
    FOREIGN KEY (platform) REFERENCES public.supported_platforms (platform) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE public.whitelist (
  platform varchar(50) NOT NULL,
  platform_user_id varchar(100) NOT NULL,
  platform_user_name varchar(100) NOT NULL,
  platform_user_login varchar(100) NOT NULL,
  user_id uuid NOT NULL,
  FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT,
  FOREIGN KEY (platform) REFERENCES public.supported_platforms (platform) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE public.blacklist (
  platform varchar(50) NOT NULL,
  platform_user_id varchar(100) NOT NULL,
  platform_user_name varchar(100) NOT NULL,
  platform_user_login varchar(100) NOT NULL,
  user_id uuid NOT NULL,
  FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT,
  FOREIGN KEY (platform) REFERENCES public.supported_platforms (platform) ON UPDATE CASCADE ON DELETE RESTRICT
);

