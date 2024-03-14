CREATE TABLE collections_new(
  id integer PRIMARY KEY,
  title text NOT NULL,
  parent_id integer REFERENCES collections(id) ON DELETE CASCADE
);

CREATE UNIQUE INDEX collections_title_parent_id_idx ON collections_new(title, parent_id);

INSERT INTO collections_new(id, title)
SELECT
  id,
  title
FROM
  collections;

DROP TABLE collections;

ALTER TABLE collections_new RENAME TO collections;

