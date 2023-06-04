CREATE TABLE users (
    id serial PRIMARY KEY,
    account varchar(100) UNIQUE NOT NULL,
    password varchar(100) NOT NULL
)