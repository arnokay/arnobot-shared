-- Create enum type "platform"
CREATE TYPE "public"."platform" AS ENUM ('twitch');
-- Modify "first_time_messages" table
ALTER TABLE "core"."first_time_messages" DROP CONSTRAINT "first_time_messages_platform_fkey", ALTER COLUMN "platform" TYPE "public"."platform" USING "platform"::"public"."platform";
-- Modify "group_chatters" table
ALTER TABLE "core"."group_chatters" DROP CONSTRAINT "group_chatters_platform_fkey", ALTER COLUMN "platform" TYPE "public"."platform" USING "platform"::"public"."platform";
-- Modify "blacklist" table
ALTER TABLE "public"."blacklist" DROP CONSTRAINT "blacklist_platform_fkey", ALTER COLUMN "platform" TYPE "public"."platform" USING "platform"::"public"."platform";
-- Modify "user_platform_accounts" table
ALTER TABLE "public"."user_platform_accounts" DROP CONSTRAINT "user_platform_accounts_platform_fkey", ALTER COLUMN "platform" TYPE "public"."platform" USING "platform"::"public"."platform";
-- Modify "whitelist" table
ALTER TABLE "public"."whitelist" DROP CONSTRAINT "whitelist_platform_fkey", ALTER COLUMN "platform" TYPE "public"."platform" USING "platform"::"public"."platform";
-- Drop "supported_platforms" table
DROP TABLE "public"."supported_platforms";
