CREATE TYPE "role" AS ENUM (
  'Admin',
  'Basic',
  'Staff',
  'Presenter'
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

CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "first_name" varchar,
  "last_name" varchar,
  "email" varchar,
  "phone_number" varchar,
  "craeted_at" timestamp,
  "last_login" timestamp,
  "role" role
);

CREATE TABLE "events" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar,
  "event_status" event_status NOT NULL DEFAULT 'inactive',
  "start" timestamp,
  "end" timestamp,
  "archived_at" timestamp,
  "craeted_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "presentations" (
  "id" SERIAL PRIMARY KEY,
  "event_id" int,
  "title" varchar,
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
  "sponsor" varchar,
  "sponsor_logo" text,
  "sponsor_description" text
);

ALTER TABLE "presentations" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");

ALTER TABLE "presenters" ADD FOREIGN KEY ("presenter_id") REFERENCES "users" ("id");

ALTER TABLE "presenters" ADD FOREIGN KEY ("presentation_id") REFERENCES "presentations" ("id");
