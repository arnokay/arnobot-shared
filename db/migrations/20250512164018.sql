-- Modify "sessions" table
ALTER TABLE "auth"."sessions" DROP CONSTRAINT "sessions_pkey", ADD PRIMARY KEY ("token");
