CREATE TABLE account (
    id SERIAL PRIMARY KEY NOT NULL,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    date_created TIMESTAMPTZ NOT NULL
);

CREATE TABLE message (
    id SERIAL PRIMARY KEY NOT NULL,
    author_id INT NOT NULL REFERENCES account(id) ON DELETE CASCADE,
    body TEXT NOT NULL,
    date_sent TIMESTAMPTZ NOT NULL
);

CREATE TABLE thread (
    id SERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE thread_member (
    thread_id INT REFERENCES thread(id) ON DELETE CASCADE,
    member_id INT REFERENCES account(id) ON DELETE CASCADE,
    PRIMARY KEY (thread_id, member_id)
);