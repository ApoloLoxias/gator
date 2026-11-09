-- name: CreateUser :one
INSERT INTO users (id, createad_at, updated_at, name)
VALUES(
  $1,
  $2,
  $3,
  $4
)
RETURNING *;
