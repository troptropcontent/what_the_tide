CREATE TABLE IF NOT EXISTS agendas (
    id INTEGER PRIMARY KEY,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    deleted_at datetime,
    type varchar NOT NULL,
    name varchar NOT NULL,
    external_id varchar,
    is_pulished boolean
)