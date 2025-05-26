-- Create enum type "webhook_status"
CREATE TYPE "twitch"."webhook_status" AS ENUM ('active', 'deactivated');
-- Modify "webhooks" table
ALTER TABLE "twitch"."webhooks" DROP CONSTRAINT "bot_id", DROP COLUMN "secret", ADD COLUMN "subscription_id" text NOT NULL, ADD COLUMN "status" "twitch"."webhook_status" NOT NULL DEFAULT 'active', ADD COLUMN "subscription_status" text NOT NULL, ADD PRIMARY KEY ("subscription_id");
