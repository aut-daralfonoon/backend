
CREATE TYPE "event_status" AS ENUM (
  'open',
  'archived',
  'inactive'
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