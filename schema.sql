CREATE TABLE todos (
  id   BIGSERIAL PRIMARY KEY,
  title text      NOT NULL,
  done boolean
);