-- Modify "bots" table
ALTER TABLE "twitch"."bots" ADD COLUMN "broadcaster_id" text NOT NULL;
-- Rename a column from "twitch_user_id" to "bot_id"
ALTER TABLE "twitch"."default_bot" RENAME COLUMN "twitch_user_id" TO "bot_id";
-- Modify "selected_bots" table
ALTER TABLE "twitch"."selected_bots" ADD COLUMN "broadcaster_id" text NOT NULL;
