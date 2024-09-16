CREATE TABLE IF NOT EXISTS ports (
    id INTEGER PRIMARY KEY,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    deleted_at datetime,
    external_id INTEGER NOT NULL,
    name varchar
)