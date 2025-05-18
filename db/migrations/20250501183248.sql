-- Add new schema named "auth"
CREATE SCHEMA "auth";
-- Add new schema named "secret"
CREATE SCHEMA "secret";
-- Add new schema named "twitch"
CREATE SCHEMA "twitch";
-- Create enum type "user_status"
CREATE TYPE "public"."user_status" AS ENUM ('ACTIVE', 'BANNED', 'DEACTIVATED', 'DELETED');
-- Create "users" table
CREATE TABLE "public"."users" ("id" integer NOT NULL GENERATED ALWAYS AS IDENTITY, "username" text NOT NULL DEFAULT '', "status" "public"."user_status" NOT NULL DEFAULT 'ACTIVE', "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create "providers" table
CREATE TABLE "auth"."providers" ("id" integer NOT NULL GENERATED ALWAYS AS IDENTITY, "user_id" integer NOT NULL, "provider" text NOT NULL, "provider_user_id" text NOT NULL, "access_token" text NOT NULL, "refresh_token" text NOT NULL, "access_type" text NOT NULL DEFAULT '', "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create enum type "bot_role"
CREATE TYPE "twitch"."bot_role" AS ENUM ('user', 'vip', 'moderator', 'broadcaster');
-- Create "default_bot" table
CREATE TABLE "twitch"."default_bot" ("main" boolean NOT NULL DEFAULT true, "twitch_user_id" text NOT NULL, PRIMARY KEY ("main"));
-- Create "default_broadcaster" table
CREATE TABLE "twitch"."default_broadcaster" ("main" boolean NOT NULL DEFAULT true, "twitch_user_id" text NOT NULL, PRIMARY KEY ("main"));
-- Create "user_bot_accounts" table
CREATE TABLE "twitch"."user_bot_accounts" ("user_id" integer NOT NULL, "twitch_user_id" text NOT NULL, "role" "twitch"."bot_role" NOT NULL DEFAULT 'user', PRIMARY KEY ("user_id", "twitch_user_id"));
-- Create "users" table
CREATE TABLE "twitch"."users" ("id" text NOT NULL, "username" text NOT NULL, "display_name" text NOT NULL, "type" text NOT NULL DEFAULT '', "broadcaster_type" text NOT NULL DEFAULT '', "profile_image_url" text NOT NULL DEFAULT '', "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "auth_provider_id" integer NULL, PRIMARY KEY ("id"));
-- Modify "providers" table
ALTER TABLE "auth"."providers" ADD CONSTRAINT "user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
-- Modify "default_bot" table
ALTER TABLE "twitch"."default_bot" ADD CONSTRAINT "twitch_user_id" FOREIGN KEY ("twitch_user_id") REFERENCES "twitch"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
-- Modify "default_broadcaster" table
ALTER TABLE "twitch"."default_broadcaster" ADD CONSTRAINT "twitch_user_id" FOREIGN KEY ("twitch_user_id") REFERENCES "twitch"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
-- Modify "user_bot_accounts" table
ALTER TABLE "twitch"."user_bot_accounts" ADD CONSTRAINT "twitch_user_id" FOREIGN KEY ("twitch_user_id") REFERENCES "twitch"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT, ADD CONSTRAINT "user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
-- Modify "users" table
ALTER TABLE "twitch"."users" ADD CONSTRAINT "auth_provider_id" FOREIGN KEY ("auth_provider_id") REFERENCES "auth"."providers" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
