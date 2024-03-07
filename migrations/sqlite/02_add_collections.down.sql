CREATE TABLE feeds_new(
  id integer PRIMARY KEY,
  url text,
  link text NOT NULL UNIQUE,
  title text NOT NULL
);

INSERT INTO feeds_new(id, url, link, title)
SELECT
  id,
  url,
  link,
  title
FROM
  feeds;

DROP TABLE feeds;

ALTER TABLE feeds_new RENAME TO feeds;

DROP TABLE collections;

