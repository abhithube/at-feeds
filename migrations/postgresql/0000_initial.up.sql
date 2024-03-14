CREATE TABLE collections(
  id serial PRIMARY KEY,
  title text NOT NULL UNIQUE
);

CREATE TABLE feeds(
  id serial PRIMARY KEY,
  url text,
  link text NOT NULL UNIQUE,
  title text NOT NULL,
  collection_id integer REFERENCES collections(id) ON DELETE CASCADE
);

CREATE TABLE entries(
  id serial PRIMARY KEY,
  link text NOT NULL UNIQUE,
  title text NOT NULL,
  published_at timestamptz,
  author text,
  content text,
  thumbnail_url text
);

CREATE TABLE feed_entries(
  feed_id integer NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
  entry_id integer NOT NULL REFERENCES entries(id) ON DELETE CASCADE,
  has_read boolean NOT NULL DEFAULT FALSE,
  PRIMARY KEY (feed_id, entry_id)
);

