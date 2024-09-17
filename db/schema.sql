CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT NOT NULL DEFAULT 'pending'
);

CREATE TABLE task_authors (
    task_id INT REFERENCES tasks(id) ON DELETE CASCADE,
    author_id INT REFERENCES authors(id) ON DELETE CASCADE,
    PRIMARY KEY (task_id, author_id)
);
