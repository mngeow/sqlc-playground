-- +goose Up
CREATE TABLE books (
  id   INTEGER PRIMARY KEY,
  name text    NOT NULL,
  author_id INTEGER REFERENCES authors(id)
);

-- +goose Down
DROP TABLE books;
