CREATE TABLE contractevents (
  seq               INTEGER         PRIMARY KEY AUTOINCREMENT,
  id                UUID            NOT NULL,
  interface_id      UUID            NOT NULL,
  namespace         VARCHAR(64)     NOT NULL,
  parent_name       VARCHAR(64)     NOT NULL,
  name              VARCHAR(64)     NOT NULL,
  params            BYTEA           NOT NULL,
  returns           BYTEA           NOT NULL
);

CREATE UNIQUE INDEX contractevents_interface_id_name ON contractevents(interface_id,name);