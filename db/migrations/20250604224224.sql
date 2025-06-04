-- Modify "sessions" table
ALTER TABLE "auth"."sessions" ALTER COLUMN "token" TYPE character varying(50);
