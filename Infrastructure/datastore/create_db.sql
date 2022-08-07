CREATE TYPE "role" AS ENUM (
  'admin',
  'inactive',
  'basic',
  'staff',
  'presenter'
);

CREATE TYPE "event_status" AS ENUM (
  'open',
  'archived',
  'inactive'
);

CREATE TYPE "session" AS ENUM (
  'closed_to_join',
  'open_to_join',
  'archived'
);

CREATE TYPE "presentation_type" AS ENUM (
  'workshop',
  'talk'
);

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "first_name" varchar,
  "last_name" varchar,
  "email" varchar NOT NULL,
  "phone_number" varchar,
  "created_at" timestamp,
  "last_login" timestamp,
  "role" role DEFAULT 'inactive'
);

CREATE TABLE "events" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar NOT NULL,
  "event_status" event_status NOT NULL DEFAULT 'inactive',
  "start" timestamp NOT NULL,
  "end" timestamp NOT NULL,
  "description" text,
  "poster" text,
  "archived_at" timestamp,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "presentations" (
  "id" SERIAL PRIMARY KEY,
  "event_id" int,
  "title" varchar NOT NULL,
  "presentation_type" presentation_type NOT NULL,
  "session_status" session NOT NULL DEFAULT 'closed_to_join',
  "capacity" int NOT NULL,
  "participant_number" int,
  "start" timestamp NOT NULL,
  "end" timestamp NOT NULL,
  "archived_at" timestamp,
  "created_at" timestamp,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "presenters" (
  "presenter_id" int,
  "presentation_id" int,
  "presenter_title" varchar,
  "presenter_description" text,
  "sponsor_name" varchar,
  "sponsor_logo" text,
  "sponsor_description" text,
  PRIMARY KEY ("presenter_id", "presentation_id")
);

CREATE TABLE "participants" (
  "user_id" int,
  "presentation_id" int
);

ALTER TABLE "presentations" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "presenters" ADD FOREIGN KEY ("presenter_id") REFERENCES "users" ("id");

ALTER TABLE "presenters" ADD FOREIGN KEY ("presentation_id") REFERENCES "presentations" ("id");

ALTER TABLE "participants" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "participants" ADD FOREIGN KEY ("presentation_id") REFERENCES "presentations" ("id");