-- Modify "selected_bots" table
ALTER TABLE "kick"."selected_bots" ADD COLUMN "enabled" boolean NOT NULL DEFAULT false;
-- Modify "selected_bots" table
ALTER TABLE "twitch"."selected_bots" ADD COLUMN "enabled" boolean NOT NULL DEFAULT false;
