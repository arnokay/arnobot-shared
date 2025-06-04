-- Add new schema named "auth"
CREATE SCHEMA "auth";
-- Add new schema named "core"
CREATE SCHEMA "core";
-- Add new schema named "secret"
CREATE SCHEMA "secret";
-- Add new schema named "twitch"
CREATE SCHEMA "twitch";
-- Create enum type "bot_role"
CREATE TYPE "twitch"."bot_role" AS ENUM ('user', 'vip', 'moderator', 'broadcaster');
-- Create "default_broadcaster" table
CREATE TABLE "twitch"."default_broadcaster" ("main" boolean NOT NULL DEFAULT true, "twitch_user_id" character varying(100) NOT NULL, PRIMARY KEY ("main"));
-- Create enum type "user_status"
CREATE TYPE "public"."user_status" AS ENUM ('active', 'banned', 'deactivated', 'deleted');
-- Create "users" table
CREATE TABLE "public"."users" ("id" uuid NOT NULL DEFAULT gen_random_uuid(), "username" character varying(50) NOT NULL DEFAULT '', "status" "public"."user_status" NOT NULL DEFAULT 'active', "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"));
-- Create enum type "session_status"
CREATE TYPE "auth"."session_status" AS ENUM ('active', 'disabled');
-- Create "default_bot" table
CREATE TABLE "twitch"."default_bot" ("main" boolean NOT NULL DEFAULT true, "bot_id" character varying(100) NOT NULL, PRIMARY KEY ("main"));
-- Create enum type "webhook_status"
CREATE TYPE "twitch"."webhook_status" AS ENUM ('active', 'deactivated');
-- Create "bots" table
CREATE TABLE "twitch"."bots" ("user_id" uuid NOT NULL, "broadcaster_id" character varying(100) NOT NULL, "bot_id" character varying(100) NOT NULL, "role" "twitch"."bot_role" NOT NULL DEFAULT 'user', PRIMARY KEY ("user_id", "bot_id"), CONSTRAINT "bots_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT);
-- Create "chatter_groups" table
CREATE TABLE "core"."chatter_groups" ("id" serial NOT NULL, "user_id" uuid NOT NULL, "name" character varying(100) NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "chatter_groups_name_user_id_key" UNIQUE ("name", "user_id"), CONSTRAINT "chatter_groups_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT);
-- Create "supported_platforms" table
CREATE TABLE "public"."supported_platforms" ("platform" character varying(50) NOT NULL, PRIMARY KEY ("platform"));
-- Create "first_time_messages" table
CREATE TABLE "core"."first_time_messages" ("platform" character varying(50) NOT NULL, "chatter_id" character varying(100) NOT NULL, "chatter_name" character varying(100) NOT NULL, "chatter_login" character varying(100) NOT NULL, "broadcaster_id" character varying(100) NOT NULL, "broadcaster_name" character varying(100) NOT NULL, "broadcaster_login" character varying(100) NOT NULL, "user_id" uuid NULL, "message" text NOT NULL, "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, CONSTRAINT "first_time_messages_platform_chatter_id_key" UNIQUE ("platform", "chatter_id"), CONSTRAINT "first_time_messages_platform_fkey" FOREIGN KEY ("platform") REFERENCES "public"."supported_platforms" ("platform") ON UPDATE CASCADE ON DELETE RESTRICT, CONSTRAINT "first_time_messages_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT);
-- Create "group_chatters" table
CREATE TABLE "core"."group_chatters" ("group_id" integer NOT NULL, "platform" character varying(50) NOT NULL, "chatter_id" character varying(100) NOT NULL, "chatter_name" character varying(100) NOT NULL, "chatter_login" character varying(100) NOT NULL, CONSTRAINT "group_chatters_group_id_platform_chatter_id_key" UNIQUE ("group_id", "platform", "chatter_id"), CONSTRAINT "group_chatters_group_id_platform_chatter_login_key" UNIQUE ("group_id", "platform", "chatter_login"), CONSTRAINT "group_chatters_group_id_fkey" FOREIGN KEY ("group_id") REFERENCES "core"."chatter_groups" ("id") ON UPDATE CASCADE ON DELETE CASCADE, CONSTRAINT "group_chatters_platform_fkey" FOREIGN KEY ("platform") REFERENCES "public"."supported_platforms" ("platform") ON UPDATE CASCADE ON DELETE RESTRICT);
-- Create "providers" table
CREATE TABLE "auth"."providers" ("id" serial NOT NULL, "user_id" uuid NOT NULL, "provider" character varying(50) NOT NULL, "provider_user_id" character varying(100) NOT NULL, "access_token" text NOT NULL, "refresh_token" text NOT NULL, "access_type" character varying(50) NOT NULL DEFAULT '', "scopes" text[] NOT NULL DEFAULT '{}', "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("id"), CONSTRAINT "providers_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE CASCADE);
-- Create "selected_bots" table
CREATE TABLE "twitch"."selected_bots" ("user_id" uuid NOT NULL, "broadcaster_id" character varying(100) NOT NULL, "bot_id" character varying(100) NOT NULL, "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("user_id"), CONSTRAINT "selected_bots_user_id_bot_id_fkey" FOREIGN KEY ("user_id", "bot_id") REFERENCES "twitch"."bots" ("user_id", "bot_id") ON UPDATE CASCADE ON DELETE RESTRICT);
-- Create "sessions" table
CREATE TABLE "auth"."sessions" ("token" uuid NOT NULL DEFAULT gen_random_uuid(), "status" "auth"."session_status" NOT NULL DEFAULT 'active', "user_id" uuid NOT NULL, "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, "last_used_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY ("token"), CONSTRAINT "sessions_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE CASCADE);
-- Create "user_commands" table
CREATE TABLE "core"."user_commands" ("user_id" uuid NOT NULL, "name" character varying(50) NOT NULL, "text" text NOT NULL, "reply" boolean NOT NULL DEFAULT false, PRIMARY KEY ("user_id", "name"), CONSTRAINT "user_commands_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT);
-- Create "user_counters" table
CREATE TABLE "core"."user_counters" ("user_id" uuid NOT NULL, "name" character varying(50) NOT NULL, "text" text NOT NULL, "count" integer NOT NULL DEFAULT 0, PRIMARY KEY ("user_id", "name"), CONSTRAINT "user_counters_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT);
-- Create "user_platform_accounts" table
CREATE TABLE "public"."user_platform_accounts" ("platform" character varying(50) NOT NULL, "platform_user_id" character varying(100) NOT NULL, "platform_user_name" character varying(100) NOT NULL, "platform_user_login" character varying(100) NOT NULL, "user_id" uuid NOT NULL, PRIMARY KEY ("platform", "platform_user_id"), CONSTRAINT "user_platform_accounts_platform_fkey" FOREIGN KEY ("platform") REFERENCES "public"."supported_platforms" ("platform") ON UPDATE CASCADE ON DELETE RESTRICT, CONSTRAINT "user_platform_accounts_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT);
-- Create "user_prefixes" table
CREATE TABLE "core"."user_prefixes" ("user_id" uuid NOT NULL, "prefix" character varying(10) NOT NULL, PRIMARY KEY ("user_id"), CONSTRAINT "user_prefixes_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT);
