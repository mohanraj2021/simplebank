CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "fullname" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashedpassword" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

-- CREATE INDEX ON "users" ("username");

ALTER TABLE "account" ADD FOREIGN KEY ("account_owner_fkey") REFERENCES "users" ("username");

ALTER TABLE "account" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner","currency");