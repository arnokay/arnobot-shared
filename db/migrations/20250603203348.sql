-- Create "supported_platforms" table
CREATE TABLE "public"."supported_platforms" ("platform" text NOT NULL, PRIMARY KEY ("platform"));
-- Create "user_platform_accounts" table
CREATE TABLE "public"."user_platform_accounts" ("platform" text NOT NULL, "platform_user_id" text NOT NULL, "platform_user_name" text NOT NULL, "platform_user_login" text NOT NULL, "user_id" integer NOT NULL, PRIMARY KEY ("platform", "platform_user_id"), CONSTRAINT "platform" FOREIGN KEY ("platform") REFERENCES "public"."supported_platforms" ("platform") ON UPDATE CASCADE ON DELETE RESTRICT, CONSTRAINT "user_id" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT);
