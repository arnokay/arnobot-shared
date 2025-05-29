-- Rename a column from "twitch_user_id" to "bot_id"
ALTER TABLE "twitch"."bots" RENAME COLUMN "twitch_user_id" TO "bot_id";
-- Rename a column from "twitch_user_id" to "bot_id"
ALTER TABLE "twitch"."selected_bots" RENAME COLUMN "twitch_user_id" TO "bot_id";
