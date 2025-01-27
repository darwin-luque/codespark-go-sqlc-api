CREATE TABLE "follow" (
  "following_user_id" uuid NOT NULL,
  "followed_user_id" uuid NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  PRIMARY KEY ("followed_user_id", "following_user_id")
);

CREATE TABLE "user" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "username" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "bio" varchar,
  "image" varchar NOT NULL,
  "password_hash" varchar NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "article" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "title" varchar NOT NULL,
  "body" text NOT NULL,
  "description" varchar,
  "slug" varchar UNIQUE NOT NULL,
  "author_id" uuid NOT NULL,
  "status" varchar,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "favorite" (
  "user_id" uuid NOT NULL,
  "article_id" uuid NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  PRIMARY KEY ("user_id", "article_id")
);

CREATE TABLE "comment" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "article_id" uuid NOT NULL,
  "author_id" uuid NOT NULL,
  "body" text NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE TABLE "tag" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "name" varchar UNIQUE NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

COMMENT ON COLUMN "article"."body" IS 'Content of the article';

COMMENT ON COLUMN "comment"."body" IS 'Content of the comment';

ALTER TABLE "article" ADD FOREIGN KEY ("author_id") REFERENCES "user" ("id");

ALTER TABLE "follow" ADD FOREIGN KEY ("following_user_id") REFERENCES "user" ("id");

ALTER TABLE "follow" ADD FOREIGN KEY ("followed_user_id") REFERENCES "user" ("id");

ALTER TABLE "favorite" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "favorite" ADD FOREIGN KEY ("article_id") REFERENCES "article" ("id");

ALTER TABLE "comment" ADD FOREIGN KEY ("article_id") REFERENCES "article" ("id");

ALTER TABLE "comment" ADD FOREIGN KEY ("author_id") REFERENCES "user" ("id");

CREATE TABLE "article_tag" (
  "article_id" uuid,
  "tag_id" uuid,
  PRIMARY KEY ("article_id", "tag_id")
);

ALTER TABLE "article_tag" ADD FOREIGN KEY ("article_id") REFERENCES "article" ("id");

ALTER TABLE "article_tag" ADD FOREIGN KEY ("tag_id") REFERENCES "tag" ("id");
