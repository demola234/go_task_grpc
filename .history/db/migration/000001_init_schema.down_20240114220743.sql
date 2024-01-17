ALTER TABLE tasks DROP CONSTRAINT tasks_user_id_fkey;

-- Drop tables in the correct order
DROP TABLE IF EXISTS tasks;

DROP TABLE IF EXISTS categories;

DROP TABLE IF EXISTS users;