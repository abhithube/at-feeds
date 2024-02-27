CREATE TABLE feeds(
  id integer PRIMARY KEY,
  url text,
  link text NOT NULL UNIQUE,
  title text NOT NULL
);

CREATE TABLE entries(
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

