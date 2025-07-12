ALTER TABLE "kick"."selected_bots" DROP CONSTRAINT "selected_bots_user_id_bot_id_fkey";
-- Modify "bots" table
ALTER TABLE "kick"."bots" ALTER COLUMN "broadcaster_id" TYPE character varying(100), ALTER COLUMN "bot_id" TYPE character varying(100);
-- Modify "default_bot" table
ALTER TABLE "kick"."default_bot" ALTER COLUMN "bot_id" TYPE character varying(100);
-- Modify "default_broadcaster" table
ALTER TABLE "kick"."default_broadcaster" ALTER COLUMN "broadcaster_id" TYPE character varying(100);
-- Modify "selected_bots" table
ALTER TABLE "kick"."selected_bots" ALTER COLUMN "broadcaster_id" TYPE character varying(100), ALTER COLUMN "bot_id" TYPE character varying(100);
ALTER TABLE "kick"."selected_bots" 
  ADD CONSTRAINT "selected_bots_user_id_bot_id_fkey" 
  FOREIGN KEY ("user_id", "bot_id") REFERENCES "kick"."bots" ("user_id", "bot_id") 
  ON UPDATE CASCADE ON DELETE RESTRICT;
