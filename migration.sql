CREATE USER 'newuser'@'localhost' IDENTIFIED BY 'password';
CREATE DATABASE chatapp;
GRANT ALL PRIVILEGES ON chatapp.* TO 'newuser'@'localhost';
GRANT ALL PRIVILEGES ON your_database.* TO 'hammad'@'localhost';
FLUSH PRIVILEGES;

-- Create the chatrooms table
CREATE TABLE IF NOT EXISTS chatrooms (
    id CHAR(36) NOT NULL DEFAULT (UUID()),
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);

-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id CHAR(36) NOT NULL DEFAULT (UUID()),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
CREATE INDEX idx_users_email ON users (email);

-- Create the messages table
CREATE TABLE IF NOT EXISTS messages (
    id CHAR(36) NOT NULL DEFAULT (UUID()),
    text TEXT NOT NULL,
    sender_id CHAR(36) NOT NULL,
    room_id CHAR(36) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (sender_id) REFERENCES users(id),
    FOREIGN KEY (room_id) REFERENCES chatrooms(id)
);
CREATE INDEX idx_messages_room_id ON messages (room_id);

CREATE TABLE IF NOT EXISTS revoked_tokens (
    token VARCHAR(255) NOT NULL,
    PRIMARY KEY (token)
);

-- Create the room_user table
CREATE TABLE IF NOT EXISTS room_user (
    room_id CHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    PRIMARY KEY (room_id, user_id),
    FOREIGN KEY (room_id) REFERENCES chatrooms(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
