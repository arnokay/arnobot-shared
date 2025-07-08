-- Modify "bots" table
ALTER TABLE "kick"."bots" ALTER COLUMN "broadcaster_id" TYPE character varying(100), ALTER COLUMN "bot_id" TYPE character varying(100);
-- Modify "default_bot" table
ALTER TABLE "kick"."default_bot" ALTER COLUMN "bot_id" TYPE character varying(100);
-- Modify "default_broadcaster" table
ALTER TABLE "kick"."default_broadcaster" ALTER COLUMN "broadcaster_id" TYPE character varying(100);
-- Modify "selected_bots" table
ALTER TABLE "kick"."selected_bots" ALTER COLUMN "broadcaster_id" TYPE character varying(100), ALTER COLUMN "bot_id" TYPE character varying(100);
