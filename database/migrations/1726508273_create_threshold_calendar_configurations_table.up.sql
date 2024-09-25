CREATE TABLE IF NOT EXISTS threshold_calendar_configurations (
    id INTEGER PRIMARY KEY,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    deleted_at datetime,
    port_id INTEGER NOT NULL,
    calendar_id INTEGER NOT NULL,
    maximum INTEGER NOT NULL,
    minimum INTEGER NOT NULL,
    FOREIGN KEY(port_id) REFERENCES ports(id),
    FOREIGN KEY(calendar_id) REFERENCES calendars(id)
)