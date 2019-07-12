CREATE TABLE usersrooms (
  user_id INTEGER REFERENCES users(id),
  room_id INTEGER REFERENCES rooms(id),
  unread BOOLEAN,
  PRIMARY KEY (user_id, room_id)
);

CREATE INDEX ON usersrooms(user_id);
CREATE INDEX ON usersrooms(room_id);
