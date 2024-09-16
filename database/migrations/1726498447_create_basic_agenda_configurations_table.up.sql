CREATE TABLE IF NOT EXISTS basic_agenda_configurations (
    id INTEGER PRIMARY KEY,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    deleted_at datetime,
    agenda_id INTEGER NOT NULL,
    port_id INTEGER NOT NULL,
    FOREIGN KEY(port_id) REFERENCES ports(id),
    FOREIGN KEY(agenda_id) REFERENCES agendas(id)
)