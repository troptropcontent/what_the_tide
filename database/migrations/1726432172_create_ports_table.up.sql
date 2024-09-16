CREATE TABLE IF NOT EXISTS ports (
    id INT PRIMARY KEY NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    deleted_at datetime NOT NULL,
    name varchar
)