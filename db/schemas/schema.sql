-- Create schemas
CREATE SCHEMA IF NOT EXISTS public;
CREATE SCHEMA IF NOT EXISTS auth;
CREATE SCHEMA IF NOT EXISTS twitch;
CREATE SCHEMA IF NOT EXISTS core;
CREATE SCHEMA IF NOT EXISTS secret;

-- Create enums
CREATE TYPE public.user_status AS ENUM ('ACTIVE', 'BANNED', 'DEACTIVATED', 'DELETED');
CREATE TYPE auth.session_status AS ENUM ('active', 'disabled');
CREATE TYPE twitch.bot_role AS ENUM ('user', 'vip', 'moderator', 'broadcaster');
CREATE TYPE twitch.webhook_status AS ENUM ('active', 'deactivated');

-- Create table: public.supported_platforms
CREATE TABLE public.supported_platforms (
  platform TEXT NOT NULL,
  PRIMARY KEY (platform)
);

-- Create table: public.users
CREATE TABLE public.users (
  id INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
  username TEXT NOT NULL DEFAULT '',
  status public.user_status NOT NULL DEFAULT 'ACTIVE',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

-- Create table: public.user_platform_accounts
CREATE TABLE public.user_platform_accounts (
  platform TEXT NOT NULL,
  platform_user_id TEXT NOT NULL,
  platform_user_name TEXT NOT NULL,
  platform_user_login TEXT NOT NULL,
  user_id INTEGER NOT NULL,
  PRIMARY KEY (platform, platform_user_id),
  FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT,
  FOREIGN KEY (platform) REFERENCES public.supported_platforms (platform) ON UPDATE CASCADE ON DELETE RESTRICT
);

-- Create table: auth.sessions
CREATE TABLE auth.sessions (
  token TEXT NOT NULL DEFAULT gen_random_uuid(),
  status auth.session_status NOT NULL DEFAULT 'active',
  user_id INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  last_used_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (token),
  FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- Create table: auth.providers
CREATE TABLE auth.providers (
  id INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY,
  user_id INTEGER NOT NULL,
  provider TEXT NOT NULL,
  provider_user_id TEXT NOT NULL,
  access_token TEXT NOT NULL,
  refresh_token TEXT NOT NULL,
  access_type TEXT NOT NULL DEFAULT '',
  scopes TEXT[] NOT NULL DEFAULT '{}',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- Create table: twitch.users
CREATE TABLE twitch.users (
  id TEXT NOT NULL,
  username TEXT NOT NULL,
  display_name TEXT NOT NULL,
  type TEXT NOT NULL DEFAULT '',
  broadcaster_type TEXT NOT NULL DEFAULT '',
  profile_image_url TEXT NOT NULL DEFAULT '',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  auth_provider_id INTEGER,
  PRIMARY KEY (id),
  FOREIGN KEY (auth_provider_id) REFERENCES auth.providers (id) ON UPDATE CASCADE ON DELETE RESTRICT
);

-- Create table: twitch.default_broadcaster
CREATE TABLE twitch.default_broadcaster (
  main BOOLEAN NOT NULL DEFAULT true,
  twitch_user_id TEXT NOT NULL,
  PRIMARY KEY (main),
  FOREIGN KEY (twitch_user_id) REFERENCES twitch.users (id) ON UPDATE CASCADE ON DELETE RESTRICT
);

-- Create table: twitch.default_bot
CREATE TABLE twitch.default_bot (
  main BOOLEAN NOT NULL DEFAULT true,
  bot_id TEXT NOT NULL,
  PRIMARY KEY (main)
);

-- Create table: twitch.bots
CREATE TABLE twitch.bots (
  user_id INTEGER NOT NULL,
  broadcaster_id TEXT NOT NULL,
  bot_id TEXT NOT NULL,
  role twitch.bot_role NOT NULL DEFAULT 'user',
  PRIMARY KEY (user_id, bot_id),
  FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT
);

-- Create table: twitch.selected_bots
CREATE TABLE twitch.selected_bots (
  user_id INTEGER NOT NULL,
  broadcaster_id TEXT NOT NULL,
  bot_id TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE (user_id),
  FOREIGN KEY (user_id, bot_id) REFERENCES twitch.bots (user_id, bot_id) ON UPDATE CASCADE ON DELETE RESTRICT
);

-- Create table: twitch.webhooks
CREATE TABLE twitch.webhooks (
  subscription_id TEXT NOT NULL,
  event TEXT NOT NULL,
  callback TEXT NOT NULL,
  user_id INTEGER NOT NULL,
  broadcaster_id TEXT NOT NULL,
  bot_id TEXT NOT NULL,
  status twitch.webhook_status NOT NULL DEFAULT 'active',
  subscription_status TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (subscription_id),
  FOREIGN KEY (broadcaster_id) REFERENCES twitch.users (id) ON UPDATE CASCADE ON DELETE RESTRICT,
  FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT,
  FOREIGN KEY (user_id, bot_id) REFERENCES twitch.bots (user_id, bot_id) ON UPDATE CASCADE ON DELETE RESTRICT
);

-- Create table: core.first_time_messages
CREATE TABLE core.first_time_messages (
  platform TEXT NOT NULL,
  platform_user_id TEXT NOT NULL,
  platform_user_name TEXT NOT NULL,
  platform_user_login TEXT NOT NULL,
  message TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  UNIQUE (platform, platform_user_id)
);
