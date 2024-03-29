CREATE TABLE
    "tasks" (
        "id" BIGSERIAL PRIMARY KEY,
        "task_name" varchar NOT NULL,
        "task_description" varchar NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now()),
        "due_date" timestamptz NOT NULL,
        "task_status" varchar NOT NULL
    );
