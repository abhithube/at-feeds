PRAGMA foreign_keys = OFF;

CREATE TABLE IF NOT EXISTS entries_new(
  id integer PRIMARY KEY,
  link text NOT NULL,
  title text NOT NULL,
  published_at text NOT NULL,
  author text,
  content text,
  thumbnail_url text,
  has_read integer NOT NULL DEFAULT 0,
  feed_id integer NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
  UNIQUE (feed_id, link)
);

INSERT INTO entries_new(id, link, title, published_at, author, content, thumbnail_url, has_read, feed_id)
SELECT
  e.id,
  e.link,
  e.title,
  e.published_at,
  e.author,
  e.content,
  e.thumbnail_url,
  fe.has_read,
  fe.feed_id
FROM
  entries e
  JOIN feed_entries fe ON fe.entry_id = e.id;

DROP TABLE entries;

ALTER TABLE entries_new RENAME TO entries;

DROP TABLE feed_entries;

PRAGMA foreign_keys = ON;

