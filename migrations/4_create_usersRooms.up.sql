CREATE TABLE usersrooms (
  user_id INTEGER REFERENCES users(id),
  room_id INTEGER REFERENCES rooms(id),
  unread BOOLEAN
);

CREATE INDEX ON messages(sender_id);
CREATE INDEX ON messages(room_id);
