CREATE TABLE items (
    id serial PRIMARY KEY,
    name TEXT NOT NULL
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

);

INSERT INTO items (name) VALUES ('laptop'), ('car'), ('bed')
