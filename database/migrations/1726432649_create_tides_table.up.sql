CREATE TABLE IF NOT EXISTS tides (
    id INTEGER PRIMARY KEY,
    port_id INT NOT NULL,
    time datetime NOT NULL,
    high boolean NOT NULL,
    coeff INT NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
    deleted_at datetime,
    FOREIGN KEY(port_id) REFERENCES ports(id)
)