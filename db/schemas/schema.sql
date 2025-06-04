-- SCHEMAS
CREATE SCHEMA IF NOT EXISTS public;

CREATE SCHEMA IF NOT EXISTS auth;

CREATE SCHEMA IF NOT EXISTS twitch;

CREATE SCHEMA IF NOT EXISTS core;

CREATE SCHEMA IF NOT EXISTS secret;

-- ENUMS
CREATE TYPE public.user_status AS ENUM (
    'ACTIVE',
    'BANNED',
    'DEACTIVATED',
    'DELETED'
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
    platform text NOT NULL,
    PRIMARY KEY (platform)
);

CREATE TABLE public.users (
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    username text NOT NULL DEFAULT '',
    status public.user_status NOT NULL DEFAULT 'ACTIVE',
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE public.user_platform_accounts (
    platform text NOT NULL,
    platform_user_id text NOT NULL,
    platform_user_name text NOT NULL,
    platform_user_login text NOT NULL,
    user_id integer NOT NULL,
    PRIMARY KEY (platform, platform_user_id),
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT,
    FOREIGN KEY (platform) REFERENCES public.supported_platforms (platform) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE auth.sessions (
    token text NOT NULL DEFAULT gen_random_uuid (),
    status auth.session_status NOT NULL DEFAULT 'active',
    user_id integer NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    last_used_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (token),
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE auth.providers (
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    user_id integer NOT NULL,
    provider text NOT NULL,
    provider_user_id text NOT NULL,
    access_token text NOT NULL,
    refresh_token text NOT NULL,
    access_type text NOT NULL DEFAULT '',
    scopes text[] NOT NULL DEFAULT '{}',
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE twitch.users (
    id text NOT NULL,
    username text NOT NULL,
    display_name text NOT NULL,
    type TEXT NOT NULL DEFAULT '',
    broadcaster_type text NOT NULL DEFAULT '',
    profile_image_url text NOT NULL DEFAULT '',
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    auth_provider_id integer,
    PRIMARY KEY (id),
    FOREIGN KEY (auth_provider_id) REFERENCES auth.providers (id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE twitch.default_broadcaster (
    main boolean NOT NULL DEFAULT TRUE,
    twitch_user_id text NOT NULL,
    PRIMARY KEY (main),
    FOREIGN KEY (twitch_user_id) REFERENCES twitch.users (id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE twitch.default_bot (
    main boolean NOT NULL DEFAULT TRUE,
    bot_id text NOT NULL,
    PRIMARY KEY (main)
);

CREATE TABLE twitch.bots (
    user_id integer NOT NULL,
    broadcaster_id text NOT NULL,
    bot_id text NOT NULL,
    ROLE twitch.bot_role NOT NULL DEFAULT 'user',
    PRIMARY KEY (user_id, bot_id),
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE twitch.selected_bots (
    user_id integer NOT NULL,
    broadcaster_id text NOT NULL,
    bot_id text NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id),
    FOREIGN KEY (user_id, bot_id) REFERENCES twitch.bots (user_id, bot_id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE twitch.webhooks (
    subscription_id text NOT NULL,
    event text NOT NULL,
    callback text NOT NULL,
    user_id integer NOT NULL,
    broadcaster_id text NOT NULL,
    bot_id text NOT NULL,
    status twitch.webhook_status NOT NULL DEFAULT 'active',
    subscription_status text NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (subscription_id),
    FOREIGN KEY (broadcaster_id) REFERENCES twitch.users (id) ON UPDATE CASCADE ON DELETE RESTRICT,
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT,
    FOREIGN KEY (user_id, bot_id) REFERENCES twitch.bots (user_id, bot_id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE core.first_time_messages (
    platform text NOT NULL,
    platform_user_id text NOT NULL,
    platform_user_name text NOT NULL,
    platform_user_login text NOT NULL,
    message text NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (platform, platform_user_id)
);
