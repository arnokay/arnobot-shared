-- atlas:import public.schema.sql


CREATE SCHEMA kick;

CREATE TABLE kick.default_broadcaster (
    main boolean PRIMARY KEY DEFAULT TRUE,
    broadcaster_id integer NOT NULL
);

CREATE TABLE kick.default_bot (
    main boolean PRIMARY KEY DEFAULT TRUE,
    bot_id integer NOT NULL
);

CREATE TABLE kick.bots (
    user_id uuid NOT NULL,
    broadcaster_id integer NOT NULL,
    bot_id integer NOT NULL,
    PRIMARY KEY (user_id, bot_id),
    FOREIGN KEY (user_id) REFERENCES public.users (id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE kick.selected_bots (
    user_id uuid NOT NULL PRIMARY KEY,
    broadcaster_id integer NOT NULL,
    bot_id integer NOT NULL,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id, bot_id) REFERENCES kick.bots (user_id, bot_id) ON UPDATE CASCADE ON DELETE RESTRICT
);

