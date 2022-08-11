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

ALTER TABLE "presenters" ADD FOREIGN KEY ("presenter_id") REFERENCES "users" ("id");

ALTER TABLE "presenters" ADD FOREIGN KEY ("presentation_id") REFERENCES "presentations" ("id");