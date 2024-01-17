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

CREATE UNIQUE INDEX "users_email_key" ON "users"("email");

CREATE UNIQUE INDEX "users_username_key" ON "users"("username");

CREATE UNIQUE INDEX "categories_category_name_key" ON "categories"("category_name");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "categories" ("category_name");

CREATE INDEX ON "tasks" ("task_name");

ALTER TABLE "tasks"
ADD
    FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "tasks"
ADD
    FOREIGN KEY ("category_id") REFERENCES "categories" ("id");