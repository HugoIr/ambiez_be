CREATE TABLE IF NOT EXISTS task (
    id bigserial PRIMARY KEY,
    title varchar NOT NULL,
    completed boolean NOT NULL,
    hour int NOT NULL,
    minute int NOT NULL
);