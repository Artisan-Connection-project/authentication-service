CREATE TYPE user_type AS ENUM('artisan', 'admin', 'customer', 'other', 'user');

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    user_type user_type,
    bio TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

INSERT INTO users (id, username, email, password_hash, full_name, user_type, bio, created_at, updated_at) VALUES
    (gen_random_uuid(), 'jdoe', 'jdoe@example.com', 'hashedpassword1', 'John Doe', 'admin', 'Admin user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'asmith', 'asmith@example.com', 'hashedpassword2', 'Alice Smith', 'user', 'Regular user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'bwilliams', 'bwilliams@example.com', 'hashedpassword3', 'Bob Williams', 'user', 'Regular user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'ccampbell', 'ccampbell@example.com', 'hashedpassword4', 'Cathy Campbell', 'user', 'Regular user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'djohnson', 'djohnson@example.com', 'hashedpassword5', 'David Johnson', 'user', 'Regular user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'eroberts', 'eroberts@example.com', 'hashedpassword6', 'Emily Roberts', 'user', 'Regular user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'fjones', 'fjones@example.com', 'hashedpassword7', 'Frank Jones', 'user', 'Regular user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'gmartin', 'gmartin@example.com', 'hashedpassword8', 'Grace Martin', 'user', 'Regular user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'hmiller', 'hmiller@example.com', 'hashedpassword9', 'Hannah Miller', 'user', 'Regular user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    (gen_random_uuid(), 'ijackson', 'ijackson@example.com', 'hashedpassword10', 'Ian Jackson', 'artisan', 'Regular user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
