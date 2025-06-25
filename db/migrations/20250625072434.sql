-- Modify "blacklist" table
ALTER TABLE "public"."blacklist" ADD CONSTRAINT "blacklist_platform_platform_user_id_key" UNIQUE ("platform", "platform_user_id"), ADD CONSTRAINT "blacklist_platform_platform_user_login_key" UNIQUE ("platform", "platform_user_login"), ADD CONSTRAINT "blacklist_platform_platform_user_name_key" UNIQUE ("platform", "platform_user_name");
-- Modify "whitelist" table
ALTER TABLE "public"."whitelist" ADD COLUMN "id" serial NOT NULL, ADD PRIMARY KEY ("id"), ADD CONSTRAINT "whitelist_platform_platform_user_id_key" UNIQUE ("platform", "platform_user_id"), ADD CONSTRAINT "whitelist_platform_platform_user_login_key" UNIQUE ("platform", "platform_user_login"), ADD CONSTRAINT "whitelist_platform_platform_user_name_key" UNIQUE ("platform", "platform_user_name");
