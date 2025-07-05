PRAGMA foreign_keys = 1;
PRAGMA journal_mode = WAL;
PRAGMA synchronous = NORMAL;

CREATE TABLE IF NOT EXISTS authors (
  id   INTEGER PRIMARY KEY,
  name TEXT    UNIQUE NOT NULL,
  bio  TEXT
) STRICT;
