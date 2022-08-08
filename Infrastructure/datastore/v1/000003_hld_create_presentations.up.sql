
CREATE TYPE "session" AS ENUM (
  'closed_to_join',
  'open_to_join',
  'archived'
);

CREATE TYPE "presentation_type" AS ENUM (
  'workshop',
  'talk'
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

ALTER TABLE "presentations" ADD FOREIGN KEY ("event_id") REFERENCES "events" ("id");