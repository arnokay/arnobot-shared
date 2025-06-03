-- Modify "providers" table
ALTER TABLE "auth"."providers" DROP CONSTRAINT "user_id";
-- Modify "sessions" table
ALTER TABLE "auth"."sessions" DROP CONSTRAINT "user_id";
-- Rename an index from "platform_to_platform_user_id" to "first_time_messages_platform_platform_user_id_key"
ALTER INDEX "core"."platform_to_platform_user_id" RENAME TO "first_time_messages_platform_platform_user_id_key";
-- Modify "user_platform_accounts" table
ALTER TABLE "public"."user_platform_accounts" DROP CONSTRAINT "platform", DROP CONSTRAINT "user_id";
-- Modify "bots" table
ALTER TABLE "twitch"."bots" DROP CONSTRAINT "user_id";
-- Modify "default_broadcaster" table
ALTER TABLE "twitch"."default_broadcaster" DROP CONSTRAINT "twitch_user_id";
-- Modify "selected_bots" table
ALTER TABLE "twitch"."selected_bots" DROP CONSTRAINT "twitch_bot";
-- Rename an index from "user_id" to "selected_bots_user_id_key"
ALTER INDEX "twitch"."user_id" RENAME TO "selected_bots_user_id_key";
-- Modify "users" table
ALTER TABLE "twitch"."users" DROP CONSTRAINT "auth_provider_id";
-- Modify "webhooks" table
ALTER TABLE "twitch"."webhooks" DROP CONSTRAINT "broadcaster_id", DROP CONSTRAINT "user_bot_account", DROP CONSTRAINT "user_id";
-- Modify "providers" table
ALTER TABLE "auth"."providers" ADD CONSTRAINT "providers_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
-- Modify "sessions" table
ALTER TABLE "auth"."sessions" ADD CONSTRAINT "sessions_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE CASCADE;
-- Modify "user_platform_accounts" table
ALTER TABLE "public"."user_platform_accounts" ADD CONSTRAINT "user_platform_accounts_platform_fkey" FOREIGN KEY ("platform") REFERENCES "public"."supported_platforms" ("platform") ON UPDATE CASCADE ON DELETE RESTRICT, ADD CONSTRAINT "user_platform_accounts_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
-- Modify "bots" table
ALTER TABLE "twitch"."bots" ADD CONSTRAINT "bots_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
-- Modify "default_broadcaster" table
ALTER TABLE "twitch"."default_broadcaster" ADD CONSTRAINT "default_broadcaster_twitch_user_id_fkey" FOREIGN KEY ("twitch_user_id") REFERENCES "twitch"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
-- Modify "selected_bots" table
ALTER TABLE "twitch"."selected_bots" ADD CONSTRAINT "selected_bots_user_id_bot_id_fkey" FOREIGN KEY ("user_id", "bot_id") REFERENCES "twitch"."bots" ("user_id", "bot_id") ON UPDATE CASCADE ON DELETE RESTRICT;
-- Modify "users" table
ALTER TABLE "twitch"."users" ADD CONSTRAINT "users_auth_provider_id_fkey" FOREIGN KEY ("auth_provider_id") REFERENCES "auth"."providers" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
-- Modify "webhooks" table
ALTER TABLE "twitch"."webhooks" ADD CONSTRAINT "webhooks_broadcaster_id_fkey" FOREIGN KEY ("broadcaster_id") REFERENCES "twitch"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT, ADD CONSTRAINT "webhooks_user_id_bot_id_fkey" FOREIGN KEY ("user_id", "bot_id") REFERENCES "twitch"."bots" ("user_id", "bot_id") ON UPDATE CASCADE ON DELETE RESTRICT, ADD CONSTRAINT "webhooks_user_id_fkey" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE CASCADE ON DELETE RESTRICT;
