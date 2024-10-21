-- Table: Users
CREATE TABLE IF NOT EXISTS Users (
    id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    createdAt TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Table: Auth
CREATE TABLE IF NOT EXISTS Auth (
    id UUID PRIMARY KEY,
    userId UUID NOT NULL,
    token TEXT NOT NULL,
    ipAddress TEXT,
    userAgent TEXT,
    expiresAt TIMESTAMPTZ,
    FOREIGN KEY (userId) REFERENCES Users(id) ON DELETE CASCADE
);

-- Index: Auth
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
