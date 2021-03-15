CREATE TABLE task_table
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task_id character varying(255),
    task_name character varying(255),
    task_done boolean not null,
    CONSTRAINT uniquetaskidconstraint UNIQUE (task_id)
);
