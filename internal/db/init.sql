CREATE TABLE IF NOT EXISTS users
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(255) NOT NULL DEFAULT 'User',
    email        VARCHAR(255) NOT NULL,
    salt         VARCHAR(20)  NOT NULL,
    password     VARCHAR(64)  NOT NULL,
    access_level SMALLINT     NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS students
(
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(255) NOT NULL,
    email   VARCHAR(255) NOT NULL,
    user_id INTEGER REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS groups
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS group_users
(
    group_id INTEGER NOT NULL REFERENCES groups (id),
    user_id  INTEGER NOT NULL REFERENCES users (id),
    PRIMARY KEY (group_id, user_id)
);

CREATE TABLE IF NOT EXISTS group_students
(
    group_id   INTEGER NOT NULL REFERENCES groups (id),
    student_id INTEGER NOT NULL REFERENCES students (id),
    PRIMARY KEY (group_id, student_id)
);

CREATE TABLE IF NOT EXISTS events
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL DEFAULT '',
    description TEXT         NOT NULL DEFAULT '',
    date        DATE         NOT NULL DEFAULT CURRENT_DATE,
    group_id    INTEGER      NOT NULL REFERENCES groups (id)
);

CREATE TABLE IF NOT EXISTS works
(
    id         SERIAL PRIMARY KEY,
    time       TIMESTAMP WITH TIME ZONE NOT NULL,
    event_id   INTEGER                  NOT NULL REFERENCES events (id),
    student_id INTEGER                  NOT NULL REFERENCES students (id)
);

CREATE TABLE IF NOT EXISTS adoptions
(
    id          SERIAL PRIMARY KEY,
    work_id     INTEGER NOT NULL REFERENCES works (id),
    path        TEXT,

    part_offset INTEGER,
    part_size   INTEGER,
    refers_to   INTEGER REFERENCES adoptions (id)
);

CREATE TABLE IF NOT EXISTS runners
(
    id    SERIAL PRIMARY KEY,
    name  VARCHAR(255) NOT NULL,
    token UUID         NOT NULL,
    tag   VARCHAR(20)  NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks
(
    id      SERIAL PRIMARY KEY,
    work_id INTEGER     NOT NULL REFERENCES works (id),
    tag     VARCHAR(20) NOT NULL,
    status  SMALLINT    NOT NULL DEFAULT 0
)