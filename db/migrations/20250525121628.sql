-- Create "bots" table
CREATE TABLE "twitch"."bots" ("user_id" integer NOT NULL, "twitch_user_id" text NOT NULL, "role" "twitch"."bot_role" NOT NULL DEFAULT 'user', PRIMARY KEY ("user_id", "twitch_user_id"), CONSTRAINT "twitch_user_id" FOREIGN KEY ("twitch_user_id") REFERENCES "twitch"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT, CONSTRAINT "user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT);
-- Create "selected_bots" table
CREATE TABLE "twitch"."selected_bots" ("user_id" integer NOT NULL, "twitch_user_id" text NOT NULL, "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, CONSTRAINT "user_id" UNIQUE ("user_id"), CONSTRAINT "twitch_bot" FOREIGN KEY ("user_id", "twitch_user_id") REFERENCES "twitch"."bots" ("user_id", "twitch_user_id") ON UPDATE CASCADE ON DELETE RESTRICT);
-- Modify "webhooks" table
ALTER TABLE "twitch"."webhooks" DROP CONSTRAINT "user_bot_account", ADD CONSTRAINT "user_bot_account" FOREIGN KEY ("user_id", "bot_id") REFERENCES "twitch"."bots" ("user_id", "twitch_user_id") ON UPDATE CASCADE ON DELETE RESTRICT;
-- Drop "bot" table
DROP TABLE "twitch"."bot";
