CREATE TYPE "user_role" AS ENUM (
  'ADMIN',
  'MOD',
  'USER'
);

CREATE TYPE "vote_type" AS ENUM (
  'UP',
  'DOWN'
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "avatar" varchar,
  "role" user_role DEFAULT 'USER',
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "subleddits" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar UNIQUE NOT NULL,
  "user_id" uuid,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "title" varchar NOT NULL,
  "url" varchar,
  "body" varchar,
  "user_id" uuid,
  "subleddit_id" uuid,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "comments" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "body" varchar NOT NULL,
  "post_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "reply_id" uuid,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "post_images" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "url" varchar NOT NULL,
  "post_id" uuid NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "post_votes" (
  "vote" vote_type NOT NULL,
  "post_id" uuid NOT NULL,
  "user_id" uuid NOT NULL
);

CREATE TABLE "comment_votes" (
  "vote" vote_type NOT NULL,
  "comment_id" uuid NOT NULL,
  "user_id" uuid NOT NULL
);

CREATE UNIQUE INDEX ON "post_votes" ("post_id", "user_id");

CREATE UNIQUE INDEX ON "comment_votes" ("comment_id", "user_id");

ALTER TABLE "subleddits" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("subleddit_id") REFERENCES "subleddits" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("reply_id") REFERENCES "comments" ("id") ON DELETE CASCADE;

ALTER TABLE "post_images" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE;

ALTER TABLE "post_votes" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE;

ALTER TABLE "post_votes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "comment_votes" ADD FOREIGN KEY ("comment_id") REFERENCES "comments" ("id") ON DELETE CASCADE;

ALTER TABLE "comment_votes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;
