-- Create "users" table
CREATE TABLE "public"."users" ("id" uuid NOT NULL, "username" character varying(50) NOT NULL, "email" character varying(100) NOT NULL, "password_hash" character varying(255) NOT NULL, "created_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP, "updated_at" timestamptz NULL DEFAULT CURRENT_TIMESTAMP, "deleted_at" timestamptz NULL, PRIMARY KEY ("id"), CONSTRAINT "users_email_key" UNIQUE ("email"), CONSTRAINT "users_username_key" UNIQUE ("username"));
-- Create index "idx_users_email" to table: "users"
CREATE INDEX "idx_users_email" ON "public"."users" ("email");
-- Create index "idx_users_username" to table: "users"
CREATE INDEX "idx_users_username" ON "public"."users" ("username");
