-- Users table
CREATE TABLE users (
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email varchar(255) NOT NULL,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    password varchar(60) NOT NULL,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now()
);

-- Tokens table
CREATE TABLE tokens (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT,
    email VARCHAR(255) NOT NULL,
    token VARCHAR(255) NOT NULL,
    token_hash BYTEA NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP,
    expiry TIMESTAMP,
    CONSTRAINT fk_tokens_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);