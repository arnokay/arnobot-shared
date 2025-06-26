-- Add new schema named "kick"
CREATE SCHEMA "kick";
-- Rename a column from "twitch_user_id" to "broadcaster_id"
ALTER TABLE "twitch"."default_broadcaster" RENAME COLUMN "twitch_user_id" TO "broadcaster_id";
-- Create "default_bot" table
CREATE TABLE "kick"."default_bot" ("main" boolean NOT NULL DEFAULT true, "bot_id" integer NOT NULL, PRIMARY KEY ("main"));
-- Create "default_broadcaster" table
CREATE TABLE "kick"."default_broadcaster" ("main" boolean NOT NULL DEFAULT true, "broadcaster_id" integer NOT NULL, PRIMARY KEY ("main"));
-- Modify "bots" table
ALTER TABLE "twitch"."bots" DROP COLUMN "role";
-- Create "bots" table
CREATE TABLE "kick"."bots" ("user_id" uuid NOT NULL, "broadcaster_id" integer NOT NULL, "bot_id" integer NOT NULL, PRIMARY KEY ("user_id", "bot_id"), CONSTRAINT "bots_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT);
-- Create "selected_bots" table
CREATE TABLE "kick"."selected_bots" ("user_id" uuid NOT NULL, "broadcaster_id" integer NOT NULL, "bot_id" integer NOT NULL, "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("user_id"), CONSTRAINT "selected_bots_user_id_bot_id_fkey" FOREIGN KEY ("user_id", "bot_id") REFERENCES "kick"."bots" ("user_id", "bot_id") ON UPDATE CASCADE ON DELETE RESTRICT);
-- Drop enum type "bot_role"
DROP TYPE "twitch"."bot_role";
