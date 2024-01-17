CREATE TABLE
    "users" (
        "id" BIGSERIAL PRIMARY KEY,
        "username" varchar NOT NULL,
        "email" varchar NOT NULL,
        "password" varchar NOT NULL,
        "role" varchar NOT NULL,
        "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
        "created_at" timestamptz NOT NULL DEFAULT (now()),
        CONSTRAINT "users_pkey" PRIMARY KEY ("id")
    );

CREATE TABLE
    "categories" (
        "id" BIGSERIAL PRIMARY KEY,
        "category_name" varchar NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now()),
        CONSTRAINT "users_pkey" PRIMARY KEY ("id")
    );

CREATE TABLE
    "tasks" (
        "id" BIGSERIAL PRIMARY KEY,
        "task_name" varchar NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now()),
        "due_date" timestamptz NOT NULL,
        "task_status" varchar NOT NULL,

CONSTRAINT "users_pkey" PRIMARY KEY ("id") );


CREATE INDEX ON "tasks" ("task_name");
