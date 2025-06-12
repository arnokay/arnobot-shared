-- atlas:import public.schema.sql


CREATE SCHEMA IF NOT EXISTS auth;

CREATE TYPE auth.session_status AS ENUM (
    'active',
    'disabled'
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

