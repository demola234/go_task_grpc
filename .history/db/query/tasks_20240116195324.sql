-- name: GetTask :one
SELECT * FROM tasks WHERE id = $1 LIMIT 1;

-- name: CreateTasks :one
INSERT INTO
    tasks (
        task_name, created_at, task_description, due_date, task_status
    )
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: ListTasks :many
SELECT * FROM tasks 


-- name: ListTasksByCategory :many
SELECT *
FROM tasks
WHERE
    task_status = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateTasks :exec
UPDATE tasks
SET
    task_name = $2,
    due_date = $3,
    task_status = $4
WHERE
    id = $1 RETURNING *;

-- name: DeleteTasks :exec
DELETE FROM tasks WHERE id = $1;