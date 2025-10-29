CREATE TABLE "users" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT (gen_random_uuid()),
  "email" varchar UNIQUE NOT NULL,
  "password_hash" varchar NOT NULL,
  "role" varchar NOT NULL DEFAULT 'reader',
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);