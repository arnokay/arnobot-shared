-- atlas:import public.schema.sql


CREATE SCHEMA IF NOT EXISTS core;

CREATE TABLE core.first_time_messages (
    platform public.platform NOT NULL,
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
    platform public.platform NOT NULL,
    chatter_id varchar(100) NOT NULL,
    chatter_name varchar(100) NOT NULL,
    chatter_login varchar(100) NOT NULL,
    UNIQUE (group_id, platform, chatter_id),
    UNIQUE (group_id, platform, chatter_login),
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

