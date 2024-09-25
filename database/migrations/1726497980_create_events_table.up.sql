CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    deleted_at datetime,
    name varchar,
    calendar_id INTEGER NOT NULL,
    FOREIGN KEY(calendar_id) REFERENCES calendars(id)
)