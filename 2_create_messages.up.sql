CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    create_time TIMESTAMP WITH TIME ZONE,
    body TEXT,
    sender_id INTEGER REFERENCES users(id),
    room_id INTEGER REFERENCES rooms(id)
);

CREATE INDEX ON messages(sender_id);
CREATE INDEX ON messages(room_id);