-- Add new schema named "core"
CREATE SCHEMA "core";
-- Create "first_time_messages" table
CREATE TABLE "core"."first_time_messages" ("platform" text NOT NULL, "platform_user_id" text NOT NULL, "platform_user_name" text NOT NULL, "platform_user_login" text NOT NULL, "message" text NOT NULL, "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, CONSTRAINT "platform_to_platform_user_id" UNIQUE ("platform", "platform_user_id"));
