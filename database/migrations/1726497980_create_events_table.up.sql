CREATE TABLE IF NOT EXISTS events (
    id INTEGER PRIMARY KEY,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    deleted_at datetime,
    name varchar,
    agenda_id INTEGER NOT NULL,
    FOREIGN KEY(agenda_id) REFERENCES agendas(id)
)