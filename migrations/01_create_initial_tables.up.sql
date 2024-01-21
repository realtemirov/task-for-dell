DROP TABLE IF EXISTS blogs CASCADE;

CREATE TABLE blogs
(
    id          SERIAL                      PRIMARY KEY,
    title       VARCHAR(255)                NOT NULL    CHECK (title <> ''),
    content     TEXT                        NOT NULL    CHECK (content <> ''),
    created_at  TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT CURRENT_TIMESTAMP 
);

CREATE TABLE news
(
    id          SERIAL                      PRIMARY KEY,
    title       VARCHAR(255)                NOT NULL    CHECK (title <> ''),
    content     TEXT                        NOT NULL    CHECK (content <> ''),
    created_at  TIMESTAMP WITH TIME ZONE    NOT NULL    DEFAULT CURRENT_TIMESTAMP 
);