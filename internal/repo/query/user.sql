-- name: GetUserByID :one
SELECT id, user_name, create_time, update_time FROM user 
WHERE id = ?;

-- name: InsertUser :execlastid
INSERT INTO user (user_name) VALUES (?);

-- name: GetUserByIDs :many
SELECT id, user_name, create_time, update_time FROM user
WHERE id IN (sqlc.slice('ids'));