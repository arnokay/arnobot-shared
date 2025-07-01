-- Modify "user_commands" table
ALTER TABLE "core"."user_commands" ADD COLUMN "created_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, ADD COLUMN "updated_at" timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP;
