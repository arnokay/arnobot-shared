# schema "public" {}
# schema "auth" {}
# schema "twitch" {}
# schema "secret" {}
# schema "core" {}
#
# table "public" "supported_platforms" {
#   schema = schema.public
#
#   column "platform" {
#     type = text
#     null = false
#   }
#
#   primary_key {
#     columns = [column.platform]
#   }
# }
#
# table "public" "user_platform_accounts" {
#   schema = schema.public
#
#   column "platform" {
#     type = text
#     null = false
#   }
#
#   column "platform_user_id" {
#     type = text
#     null = false
#   }
#
#   column "platform_user_name" {
#     type = text
#     null = false
#   }
#
#   column "platform_user_login" {
#     type = text
#     null = false
#   }
#
#   column "user_id" {
#     type = int
#     null = false
#   }
#
#   primary_key {
#     columns = [column.platform, column.platform_user_id]
#   }
#
#   foreign_key "user_id" {
#     columns = [column.user_id]
#     ref_columns = [table.public.users.column.id]
#     on_update = CASCADE
#     on_delete = RESTRICT
#   }
#
#   foreign_key "platform" {
#     columns = [column.platform]
#     ref_columns = [table.public.supported_platforms.column.platform]
#     on_update = CASCADE
#     on_delete = RESTRICT
#   }
# }
#
# enum "user_status" {
#   schema = schema.public
#   values = [
#   "ACTIVE",
#   "BANNED",
#   "DEACTIVATED",
#   "DELETED"
#   ]
# }
#
# table "public" "users" {
#   schema = schema.public
#
#   column "id" {
#     type = int
#     null = false
#     identity { generated = ALWAYS }
#   }
#
#   column "username" {
#     type = text
#     null = false
#     default = ""
#   }
#
#   column "status" {
#     type = enum.user_status
#     null = false
#     default = "ACTIVE"
#   }
#
#   column "created_at" {
#     type = timestamp
#     null = false
#     default = sql("CURRENT_TIMESTAMP")
#   }
#
#   column "updated_at" {
#     type = timestamp
#     null = false
#     default = sql("CURRENT_TIMESTAMP")
#   }
#
#   primary_key {
#     columns = [column.id]
#   }
# }
#
# enum "auth" "session_status" {
#   schema = schema.auth
#   values = [
#     "active",
#     "disabled"
#   ]
# }
#
# table "auth" "sessions" {
#   schema = schema.auth
#
#   column "token" {
#     type = text
#     null = false
#     default = sql("gen_random_uuid()")
#   }
#
#   column "status" {
#     type = enum.auth.session_status
#     null = false
#     default = "active"
#   }
#
#   column "user_id" {
#     type = int
#     null = false
#   }
#
#   column "created_at" {
#     type = timestamp
#     null = false
#     default = sql("CURRENT_TIMESTAMP")
#   }
#
#   column "last_used_at" {
#     type = timestamp
#     null = false
#     default = sql("CURRENT_TIMESTAMP")
#   }
#   
#   foreign_key "user_id" {
#     columns = [column.user_id]
#     ref_columns = [table.public.users.column.id]
#     on_update = CASCADE
#     on_delete = CASCADE
#   }
#
#   primary_key {
#     columns = [column.token]
#   }
# }
#
# table "auth" "providers" {
#   schema = schema.auth
#
#   column "id" {
#     type = int
#     null = false
#     identity { generated = ALWAYS }
#   }
#
#   column "user_id" {
#     type = int
#     null = false
#   }
#
#   column "provider" {
#     type = text
#     null = false
#   }
#
#   column "provider_user_id" {
#     type = text
#     null = false
#   }
#
#   column "access_token" {
#     type = text
#     null = false
#   }
#
#   column "refresh_token" {
#     type = text
#     null = false
#   }
#
#   column "access_type" {
#     type = text
#     null = false
#     default = ""
#   }
#
#   column "scopes" {
#     type = sql("text[]")
#     null = false
#     default = "{}"
#   }
#
#   column "created_at" {
#     type = timestamp
#     null = false
#     default = sql("CURRENT_TIMESTAMP")
#   }
#
#   column "updated_at" {
#     type = timestamp
#     null = false
#     default = sql("CURRENT_TIMESTAMP")
#   }
#
#   foreign_key "user_id" {
#     columns = [column.user_id]
#     ref_columns = [table.public.users.column.id]
#     on_update = CASCADE
#     on_delete = CASCADE
#   }
#
#   primary_key {
#     columns = [column.id]
#   }
# }
#
# table "twitch" "users" {
#   schema = schema.twitch
#
#   column "id" {
#     type = text
#     null = false
#   }
#
#   column "username" {
#     type = text
#     null = false
#   }
#
#   column "display_name" {
#     type = text
#     null = false
#   }
#
#   column "type" {
#     type = text
#     null = false
#     default = ""
#   }
#
#   column "broadcaster_type" {
#     type = text
#     null = false
#     default = ""
#   }
#
#   column "profile_image_url" {
#     type = text
#     null = false
#     default = ""
#   }
#
#   column "created_at" {
#     type = timestamp
#     null = false
#     default = sql("CURRENT_TIMESTAMP")
#   }
#
#   column "auth_provider_id" {
#     type = int
#     null = true
#   }
#
#   primary_key {
#     columns = [column.id]
#   }
#
#   foreign_key "auth_provider_id" {
#     columns = [column.auth_provider_id]
#     ref_columns = [table.auth.providers.column.id]
#     on_update = CASCADE
#     on_delete = RESTRICT 
#   }
# }
#
# table "twitch" "default_broadcaster" {
#   schema = schema.twitch
#
#   column "main" {
#     type = boolean
#     null = false
#     default = true
#   }
#
#   column "twitch_user_id" {
#     type = text
#     null = false
#   }
#   
#   primary_key {
#     columns = [column.main]
#   }
#
#   foreign_key "twitch_user_id" {
#     columns = [column.twitch_user_id]
#     ref_columns = [table.twitch.users.column.id]
#     on_update = CASCADE
#     on_delete = RESTRICT
#   }
# }
#
# table "twitch" "default_bot" {
#   schema = schema.twitch
#
#   column "main" {
#     type = boolean
#     null = false
#     default = true
#   }
#
#   column "bot_id" {
#     type = text
#     null = false
#   }
#
#   primary_key {
#     columns = [column.main]
#   }
# }
#
# enum "twitch" "bot_role" { 
#   schema = schema.twitch
#   values = [
#   "user",
#   "vip",
#   "moderator",
#   "broadcaster"
#   ]
# }
#
# table "twitch" "selected_bots" {
#   schema = schema.twitch
#
#   column "user_id" {
#     type = int
#     null = false
#   }
#   
#   column "broadcaster_id" {
#     type = text
#     null = false
#   }
#
#   column "bot_id" {
#     type = text
#     null = false
#   }
#
#   column "created_at" {
#     type = timestamp
#     null = false
#     default = sql("CURRENT_TIMESTAMP")
#   }
#
#   column "updated_at" {
#     type = timestamp
#     null = false
#     default = sql("CURRENT_TIMESTAMP")
#   }
#
#   unique "user_id" {
#     columns = [column.user_id]
#   }
#
#   foreign_key "twitch_bot" { 
#     columns = [column.user_id, column.bot_id]
#     ref_columns = [table.twitch.bots.column.user_id, table.twitch.bots.column.bot_id]
#     on_update = CASCADE
#     on_delete = RESTRICT
#   }
# }
#
# table "twitch" "bots" {
#   schema = schema.twitch
#
#   column "user_id" {
#     type = int
#     null = false
#   }
#
#   column "broadcaster_id" {
#     type = text
#     null = false
#   }
#
#   column "bot_id" {
#     type = text
#     null = false
#   }
#
#   column "role" {
#     type = enum.twitch.bot_role
#     null = false
#     default = "user"
#   }
#
#   primary_key {
#     columns = [column.user_id, column.bot_id]
#   }
#
#   foreign_key "user_id" {
#     columns = [column.user_id]
#     ref_columns = [table.public.users.column.id]
#     on_update = CASCADE
#     on_delete = RESTRICT
#   }
# }
#
# enum "twitch" "webhook_status" {
#   schema = schema.twitch
#   values = [
#   "active",
#   "deactivated" 
#   ]
# }
#
# table "twitch" "webhooks" {
#   schema = schema.twitch
#
#   column "subscription_id" {
#     type = text
#     null = false
#   }
#
#   column "event" {
#     type = text
#     null = false
#   }
#
#   column "callback" {
#     type = text
#     null = false
#   }
#
#   column "user_id" {
#     type = int
#     null = false
#   }
#
#   column "broadcaster_id" {
#     type = text
#     null = false
#   }
#
#   column "bot_id" {
#     type = text
#     null = false
#   }
#
#   column "status" {
#     type = enum.twitch.webhook_status
#     null = false
#     default = "active"
#   }
#
#   column "subscription_status" {
#     type = text
#     null = false
#   }
#
#   column "created_at" {
#     type = timestamp
#     null = false
#     default = sql("CURRENT_TIMESTAMP")
#   }
#
#   column "updated_at" {
#     type = timestamp
#     null = false
#     default = sql("CURRENT_TIMESTAMP")
#   }
#
#   primary_key {
#     columns = [column.subscription_id]
#   }
#
#   foreign_key "broadcaster_id" {
#     columns = [column.broadcaster_id]
#     ref_columns = [table.twitch.users.column.id]
#     on_update = CASCADE
#     on_delete = RESTRICT
#   }
#
#   foreign_key "user_id" {
#     columns = [column.user_id]
#     ref_columns = [table.public.users.column.id]
#     on_update = CASCADE
#     on_delete = RESTRICT
#   }
#
#   foreign_key "user_bot_account" {
#     columns = [column.user_id, column.bot_id]
#     ref_columns = [table.twitch.bots.column.user_id, table.twitch.bots.column.bot_id]
#     on_update = CASCADE
#     on_delete = RESTRICT
#   }
# }
#
# table "core" "first_time_messages" {
#   schema = schema.core
#
#   column "platform" {
#     type = text
#     null = false
#   }
#
#   column "platform_user_id" {
#     type = text
#     null = false
#   }
#
#   column "platform_user_name" {
#     type = text
#     null = false
#   }
#
#   column "platform_user_login" {
#     type = text
#     null = false
#   }
#
#   column "message" {
#     type = text
#     null = false
#   }
#
#   column "created_at" {
#     type = timestamp
#     null = false
#     default = sql("CURRENT_TIMESTAMP")
#   }
#
#   unique "platform_to_platform_user_id" {
#     columns = [column.platform, column.platform_user_id]
#   }
# }
