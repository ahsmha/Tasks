# Tasks Backend
Backend of a Tasks application

## Features
- Basic CRUD functionality

## SQL queries
```SQL
CREATE TABLE users (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    role VARCHAR(255),
    created DATETIME,
    updated DATETIME
);

CREATE TABLE tasks (
    id INT PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    due_date DATETIME,
    status VARCHAR(255),
    creator_id INT,
    created DATETIME,
    updated DATETIME,
    FOREIGN KEY (creator_id) REFERENCES users(id)
);

CREATE TABLE task_assignee_mapping (
    task_id INT,
    assignee_id INT,
    FOREIGN KEY (task_id) REFERENCES tasks(id),
    FOREIGN KEY (assignee_id) REFERENCES users(id)
);
```