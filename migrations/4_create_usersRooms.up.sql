CREATE TABLE usersRooms (
  user_id INTEGER REFERENCES users(id),
  room_id INTEGER REFERENCES rooms(id),
  unread BIT
);

CREATE INDEX ON messages(sender_id);
CREATE INDEX ON messages(room_id);
