-- name: CreateUser :one
INSERT INTO users as u(username,password,full_name,auth_code,email) 
VALUES ($1, $2, $3, $4,$5) 
RETURNING *;

-- name: GetUser :one
SELECT * from users WHERE username = $1 LIMIT 1;



-- name: GetUsers :many
select * from users;

-- name: AuthCodeUser :exec
UPDATE users u
SET auth_code= $1
where u.id =$2;

