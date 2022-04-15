CREATE TABLE items (
    id serial PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO items (name) VALUES ('laptop'), ('car'), ('bed')
