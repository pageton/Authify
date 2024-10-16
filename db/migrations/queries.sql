-- name: GetUser :one
SELECT *
FROM Users
WHERE id = ?
    OR username = ?
LIMIT 1;
-- name: GetUsers :many
SELECT *
FROM Users
ORDER BY username;
-- name: CreateUser :one
INSERT INTO Users (id, username, password)
VALUES (?, ?, ?)
RETURNING *;
-- name: UpdateUser :exec
UPDATE Users
SET username = ?,
    password = ?,
    updatedAt = CURRENT_TIMESTAMP
WHERE id = ?
    OR username = ?;
-- name: DeleteUser :exec
DELETE FROM Users
WHERE id = ?
    OR username = ?;
-- name: GetAuthToken :one
SELECT a.*
FROM Auth a
JOIN Users u ON a.userId = u.id
WHERE a.id = ? OR u.username = ?
LIMIT 1;
-- name: CreateAuthToken :exec
INSERT OR REPLACE INTO Auth (id, userId, token, expiresAt, ipAddress, userAgent)
VALUES (?, ?, ?, ?, ?, ?);
-- name: DeleteAuthToken :exec
DELETE FROM Auth
WHERE id = ?;