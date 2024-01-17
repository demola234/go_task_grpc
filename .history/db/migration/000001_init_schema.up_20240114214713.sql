CREATE TABLE
    "tasks" (
        "id" BIGSERIAL PRIMARY KEY,
        "task_name" varchar NOT NULL,
        "user_id" bigint NOT NULL,
        "category_id" bigint NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now()),
        "due_date" timestamptz NOT NULL,
        "task_status" varchar NOT NULL,
    );

ALTER TABLE "tasks"