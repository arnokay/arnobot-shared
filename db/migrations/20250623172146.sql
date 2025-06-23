-- Modify "blacklist" table
ALTER TABLE "public"."blacklist" ALTER COLUMN "platform_user_id" DROP NOT NULL, ALTER COLUMN "platform_user_name" DROP NOT NULL, ALTER COLUMN "platform_user_login" DROP NOT NULL, ALTER COLUMN "user_id" DROP NOT NULL;
-- Modify "whitelist" table
ALTER TABLE "public"."whitelist" ALTER COLUMN "platform_user_id" DROP NOT NULL, ALTER COLUMN "platform_user_name" DROP NOT NULL, ALTER COLUMN "platform_user_login" DROP NOT NULL, ALTER COLUMN "user_id" DROP NOT NULL;
