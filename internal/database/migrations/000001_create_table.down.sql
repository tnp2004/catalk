BEGIN;

DROP TRIGGER IF EXISTS set_updated_at_timestamp_users_table ON "users";
DROP TRIGGER IF EXISTS set_updated_at_timestamp_providers_table ON "providers";
DROP TRIGGER IF EXISTS set_updated_at_timestamp_chats_table ON "chats";

DROP FUNCTION IF EXISTS set_updated_at_column;

DROP TABLE IF EXISTS "users" CASCADE;
DROP TABLE IF EXISTS "providers" CASCADE;
DROP TABLE IF EXISTS "messages" CASCADE;
DROP TABLE IF EXISTS "chats" CASCADE;

COMMIT;