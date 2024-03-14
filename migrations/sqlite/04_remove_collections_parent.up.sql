CREATE TABLE collections_new(
  id integer PRIMARY KEY,
  title text NOT NULL UNIQUE
);

INSERT INTO collections_new(id, title)
SELECT
  id,
  title
FROM
  collections;

DROP TABLE collections;

ALTER TABLE collections_new RENAME TO collections;

