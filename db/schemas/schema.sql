CREATE SCHEMA IF NOT EXISTS public;

CREATE SCHEMA IF NOT EXISTS auth;

CREATE SCHEMA IF NOT EXISTS twitch;

CREATE SCHEMA IF NOT EXISTS core;

CREATE SCHEMA IF NOT EXISTS secret;

CREATE TYPE public.user_status AS ENUM (
    'active',
    'banned',
    'deactivated',
    'deleted'
);

CREATE TYPE auth.session_status AS ENUM (
    'active',
    'disabled'
);

CREATE TYPE twitch.bot_role AS ENUM (
    'user',
    'vip',
    'moderator',
    'broadcaster'
);

CREATE TYPE twitch.webhook_status AS ENUM (
    'active',
    'deactivated'
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

CREATE TABLE auth.sessions (
    token varchar(50) PRIMARY KEY DEFAULT gen_random_uuid (),
    status auth.session_status NOT NULL DEFAULT 'active',
    user_id uuid NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_used_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE auth.providers (
    id serial PRIMARY KEY,
    user_id uuid NOT NULL,
    provider varchar(50) NOT NULL,
    provider_user_id varchar(100) NOT NULL,
    access_token text NOT NULL,
    refresh_token text NOT NULL,
    access_type varchar(50) NOT NULL DEFAULT '',
    scopes text[] NOT NULL DEFAULT '{}',
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE twitch.default_broadcaster (
    main boolean PRIMARY KEY DEFAULT TRUE,
    twitch_user_id varchar(100) NOT NULL
);

CREATE TABLE twitch.default_bot (
    main boolean PRIMARY KEY DEFAULT TRUE,
    bot_id varchar(100) NOT NULL
);

CREATE TABLE twitch.bots (
    user_id uuid NOT NULL,
    broadcaster_id varchar(100) NOT NULL,
    bot_id varchar(100) NOT NULL,
    ROLE twitch.bot_role NOT NULL DEFAULT 'user',
    PRIMARY KEY (user_id, bot_id),
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE twitch.selected_bots (
    user_id uuid NOT NULL PRIMARY KEY,
    broadcaster_id varchar(100) NOT NULL,
    bot_id varchar(100) NOT NULL,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id, bot_id) REFERENCES twitch.bots (user_id, bot_id) ON UPDATE CASCADE ON DELETE RESTRICT
);

-- CREATE TABLE twitch.webhooks (
--     subscription_id text NOT NULL,
--     event text NOT NULL,
--     callback text NOT NULL,
--     user_id integer NOT NULL,
--     broadcaster_id text NOT NULL,
--     bot_id text NOT NULL,
--     status twitch.webhook_status NOT NULL DEFAULT 'active',
--     subscription_status text NOT NULL,
--     created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     PRIMARY KEY (subscription_id),
--     FOREIGN KEY (broadcaster_id) REFERENCES twitch.users (id) ON UPDATE CASCADE ON DELETE RESTRICT,
--     FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT,
--     FOREIGN KEY (user_id, bot_id) REFERENCES twitch.bots (user_id, bot_id) ON UPDATE CASCADE ON DELETE RESTRICT
-- );
CREATE TABLE core.first_time_messages (
    platform varchar(50) NOT NULL,
    chatter_id varchar(100) NOT NULL,
    chatter_name varchar(100) NOT NULL,
    chatter_login varchar(100) NOT NULL,
    broadcaster_id varchar(100) NOT NULL,
    broadcaster_name varchar(100) NOT NULL,
    broadcaster_login varchar(100) NOT NULL,
    user_id uuid,
    message text NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT,
    FOREIGN KEY (platform) REFERENCES public.supported_platforms (platform) ON UPDATE CASCADE ON DELETE RESTRICT,
    UNIQUE (platform, chatter_id)
);

CREATE TABLE core.chatter_groups (
    id serial PRIMARY KEY,
    user_id uuid NOT NULL,
    name varchar(100) NOT NULL,
    UNIQUE (name, user_id),
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE core.group_chatters (
    group_id integer NOT NULL,
    platform varchar(50) NOT NULL,
    chatter_id varchar(100) NOT NULL,
    chatter_name varchar(100) NOT NULL,
    chatter_login varchar(100) NOT NULL,
    UNIQUE (group_id, platform, chatter_id),
    UNIQUE (group_id, platform, chatter_login),
    FOREIGN KEY (platform) REFERENCES public.supported_platforms (platform) ON UPDATE CASCADE ON DELETE RESTRICT,
    FOREIGN KEY (group_id) REFERENCES core.chatter_groups (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE core.user_prefixes (
    user_id uuid PRIMARY KEY,
    prefix varchar(10) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE core.user_commands (
    user_id uuid NOT NULL,
    name varchar(50) NOT NULL,
    text text NOT NULL,
    reply boolean NOT NULL DEFAULT FALSE,
    PRIMARY KEY (user_id, name),
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE core.user_counters (
    user_id uuid NOT NULL,
    name varchar(50) NOT NULL,
    text text NOT NULL,
    count integer NOT NULL DEFAULT 0,
    PRIMARY KEY (user_id, name),
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT
);

