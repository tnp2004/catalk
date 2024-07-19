BEGIN;


INSERT INTO "providers" ("name") 
VALUES ('google'),('discord');

INSERT INTO "users" ("id", "email", "username", "picture_url", "provider_id") 
VALUES ('38ff22c8-3adc-46af-8208-24f4b545b895', 'test@gmail.com', 'testJa', 'test.img', '1'), 
('b3319923-7d7d-4f2c-ba47-e96fed21f1cf', 'rainy555@gmail.com', 'rainy', 'rainy.img', '2');

INSERT INTO "chats" ("id", "user_id", "title") 
VALUES ('065742ea-7e9d-40ab-b428-cd6fb686d8dd', '38ff22c8-3adc-46af-8208-24f4b545b895', 'How to write SQL?'), 
('207c7abb-bf25-4290-a3af-0ae12f836834', 'b3319923-7d7d-4f2c-ba47-e96fed21f1cf', 'How to play Elden ring');

INSERT INTO "messages" ("chat_id", "user_message", "model_message") 
VALUES ('065742ea-7e9d-40ab-b428-cd6fb686d8dd', 'Hello', 'Hi human!'), 
('207c7abb-bf25-4290-a3af-0ae12f836834', 'what your name?', 'My name is Boba');

COMMIT;