CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    userName VARCHAR(255),  
    email VARCHAR(255),
    jwt_token VARCHAR(255),
    password_digest BYTEA
);

CREATE UNIQUE INDEX ON users(userName);
CREATE UNIQUE INDEX on users(email);
CREATE UNIQUE INDEX ON users(jwt_token);