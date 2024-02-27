PRAGMA foreign_keys = OFF;

CREATE TABLE IF NOT EXISTS entries_new(
  id integer PRIMARY KEY,
  link text NOT NULL UNIQUE,
  title text NOT NULL,
  published_at text NOT NULL,
  author text,
  content text,
  thumbnail_url text
);

INSERT INTO entries_new(id, link, title, published_at, author, content, thumbnail_url)
SELECT
  id,
  link,
  title,
  published_at,
  author,
  content,
  thumbnail_url
FROM
  entries;

CREATE TABLE IF NOT EXISTS feed_entries(
  feed_id integer NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
  entry_id integer NOT NULL REFERENCES entries(id) ON DELETE CASCADE,
  has_read integer NOT NULL DEFAULT 0,
  PRIMARY KEY (feed_id, entry_id)
);

INSERT INTO feed_entries(feed_id, entry_id, has_read)
SELECT
  feed_id,
  id,
  has_read
FROM
  entries;

DROP TABLE entries;

ALTER TABLE entries_new RENAME TO entries;

PRAGMA foreign_keys = ON;

