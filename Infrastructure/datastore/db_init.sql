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
  "email" varchar,
  "phone_number" varchar,
  "craeted_at" timestamp,
  "last_login" timestamp,
  "role" role DEFAULT 'inactive'
);

CREATE TABLE "events" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar,
  "event_status" event_status NOT NULL DEFAULT 'inactive',
  "start" timestamp,
  "end" timestamp,
  "description" text,
  "poster" text,
  "archived_at" timestamp,
  "craeted_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "presentations" (
  "id" SERIAL PRIMARY KEY,
  "event_id" int,
  "title" varchar,
  "presentation_type" presentation_type,
  "session_status" session,
  "capacity" int,
  "participant_number" int,
  "start" timestamp,
  "end" timestamp,
  "archived_at" timestamp,
  "craeted_at" timestamp,
  "updated_at" timestamp
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
