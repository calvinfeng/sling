CREATE TABLE rooms (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    room_type BIT
);

CREATE UNIQUE INDEX ON rooms(id);
