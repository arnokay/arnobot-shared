-- atlas:import public.schema.sql


CREATE SCHEMA IF NOT EXISTS twitch;

CREATE TYPE twitch.bot_role AS ENUM (
    'user',
    'vip',
    'moderator',
    'broadcaster'
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
