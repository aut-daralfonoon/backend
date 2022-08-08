CREATE TYPE "role" AS ENUM (
  'admin',
  'inactive',
  'basic',
  'staff',
  'presenter'
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
