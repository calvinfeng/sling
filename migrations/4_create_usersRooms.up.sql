CREATE TABLE usersrooms (
  user_id INTEGER REFERENCES users(id),
  room_id INTEGER REFERENCES rooms(id),
  unread BOOLEAN
);

CREATE INDEX ON usersrooms(user_id);
CREATE INDEX ON usersrooms(room_id);
