CREATE TABLE subscriptions (
    id INTEGER PRIMARY KEY,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    deleted_at datetime,
    agenda_id INTEGER NOT NULL,
    email varchar NOT NULL,
    is_published boolean NOT NULL,
    external_id varchar,
    FOREIGN KEY(agenda_id) REFERENCES agendas(id)
)