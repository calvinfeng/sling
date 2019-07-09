CREATE TABLE userroom (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    room_id INTEGER REFERENCES rooms(id),
    hasUnread boolean
);

CREATE INDEX ON userroom(user_id);
CREATE INDEX ON userroom(room_id);