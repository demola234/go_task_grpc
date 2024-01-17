
CREATE TABLE
    "tasks" (
        "id" BIGSERIAL PRIMARY KEY,
        "task_name" varchar NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now()),
        "due_date" timestamptz NOT NULL,
        "task_status" varchar NOT NULL,

CONSTRAINT "users_pkey" PRIMARY KEY ("id") );


CREATE INDEX ON "tasks" ("task_name");
