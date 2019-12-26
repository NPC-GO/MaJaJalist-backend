CREATE TABLE IF NOT EXISTS todos
(
    id          VARCHAR(30) PRIMARY KEY,
    textContent VARCHAR(100),
    authorID    VARCHAR(30) NOT NULL,
    finished    boolean     NOT NULL DEFAULT false,
    complete    boolean     NOT NULL DEFAULT false,
    deleted     boolean     NOT NULL DEFAULT false
);