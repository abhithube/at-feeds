CREATE TABLE collections_new(
  id integer PRIMARY KEY,
  title text UNIQUE NOT NULL,
  parent_id integer REFERENCES collections(id) ON DELETE CASCADE
);

INSERT INTO collections_new(id, title, parent_id)
SELECT
  id,
  title,
  parent_id
FROM
  collections;

DROP TABLE collections;

ALTER TABLE collections_new RENAME TO collections;

