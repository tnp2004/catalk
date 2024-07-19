BEGIN;

SET TIME ZONE 'Asia/Bangkok';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION set_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  return NEW;
END;
$$ language 'plpgsql';

CREATE TABLE "users" (
  "id" VARCHAR PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "email" VARCHAR UNIQUE NOT NULL,
  "username" VARCHAR NOT NULL,
  "picture_url" VARCHAR NOT NULL,
  "provider_id" SERIAL NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMP NOT NULL DEFAULT (now())
);

CREATE TABLE "providers" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMP NOT NULL DEFAULT (now())
);

CREATE TABLE "messages" (
  "id" VARCHAR PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "chat_id" VARCHAR NOT NULL,
  "user_message" VARCHAR NOT NULL,
  "model_message" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (now())
);

CREATE TABLE "chats" (
  "id" VARCHAR PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
  "user_id" VARCHAR NOT NULL,
  "title" VARCHAR NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT (now()),
  "updated_at" TIMESTAMP NOT NULL DEFAULT (now())
);

ALTER TABLE "users" ADD FOREIGN KEY ("provider_id") REFERENCES "providers" ("id");
ALTER TABLE "messages" ADD FOREIGN KEY ("chat_id") REFERENCES "chats" ("id");
ALTER TABLE "chats" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE TRIGGER set_updated_at_timestamp_users_table BEFORE UPDATE ON "users" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_providers_table BEFORE UPDATE ON "providers" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();
CREATE TRIGGER set_updated_at_timestamp_chats_table BEFORE UPDATE ON "chats" FOR EACH ROW EXECUTE PROCEDURE set_updated_at_column();

COMMIT;