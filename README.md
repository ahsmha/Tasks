# Tasks Backend
Backend of a Tasks application

## Features
- Basic CRUD functionality

## SQL queries
```SQL

CREATE TABLE tasks.users (
    id INT AUTO_INCREMENT,
    name VARCHAR(255),
    role VARCHAR(255),
    created DATETIME,
    updated DATETIME,
	PRIMARY KEY(id)
);

CREATE TABLE tasks.tasks (
    id INT AUTO_INCREMENT,
    title VARCHAR(50) NOT NULL,
    due_date DATETIME,
    status VARCHAR(255),
    creator_id INT,
    created DATETIME,
    updated DATETIME,
	PRIMARY KEY(id),
    FOREIGN KEY (creator_id) REFERENCES users(id)
);

CREATE TABLE tasks.task_assignee_mapping (
    task_id INT,
    assignee_id INT,
    FOREIGN KEY (task_id) REFERENCES tasks(id),
    FOREIGN KEY (assignee_id) REFERENCES users(id)
);


INSERT INTO tasks.users (name, role, created, updated) VALUES
  ('John Doe', 'lead', '2023-07-12 02:11:17', '2023-07-12 02:11:17'),
  ('Jane Doe', 'subordinate', '2023-07-12 02:11:17', '2023-07-12 02:11:17'),
  ('Mary Smith', 'lead', '2023-07-12 02:11:17', '2023-07-12 02:11:17'),
  ('Joh Doe', 'subordinate', '2023-07-12 02:11:17', '2023-07-12 02:11:17');
```