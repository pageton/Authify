-- Table: Users
CREATE TABLE IF NOT EXISTS Users (
    id TEXT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table: Auth
CREATE TABLE IF NOT EXISTS Auth (
    id TEXT PRIMARY KEY,
    userId TEXT NOT NULL,
    token TEXT NOT NULL,
    ipAddress TEXT,
    userAgent TEXT,
    expiresAt TIMESTAMP,
    FOREIGN KEY (userId) REFERENCES Users(id) ON DELETE CASCADE
);

-- Index: Auth
CREATE INDEX idx_userId ON Auth (userId);
