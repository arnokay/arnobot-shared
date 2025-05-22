-- Modify "providers" table
ALTER TABLE "auth"."providers" ADD COLUMN "scopes" text[] NOT NULL DEFAULT '{}';
