-- name: GetUser :one
SELECT *
FROM Users
WHERE id = $1
    OR username = $2
LIMIT 1;
-- name: GetUsers :many
SELECT *
FROM Users
ORDER BY username;
-- name: CreateUser :one
INSERT INTO Users (id, username, password)
VALUES ($1, $2, $3)
RETURNING *;
-- name: UpdateUser :exec
UPDATE Users
SET username = $1,
    password = $2,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = $3
    OR username = $4;
-- name: DeleteUser :exec
DELETE FROM Users
WHERE id = $1
    OR username = $2;
-- name: GetAuthToken :one
SELECT a.*
FROM Auth a
JOIN Users u ON a.userId = u.id
WHERE a.id = $1 OR u.username = $2
LIMIT 1;
-- name: CreateAuthToken :exec
INSERT INTO Auth (id, userId, token, expiresAt, ipAddress, userAgent)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (id) DO UPDATE
SET token = EXCLUDED.token,
    expiresAt = EXCLUDED.expiresAt,
    ipAddress = EXCLUDED.ipAddress,
    userAgent = EXCLUDED.userAgent;
-- name: DeleteAuthToken :exec
DELETE FROM Auth
WHERE id = $1;
