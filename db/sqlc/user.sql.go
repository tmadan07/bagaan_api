// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const authCodeUser = `-- name: AuthCodeUser :exec
UPDATE users u
SET auth_code= $1
where u.id =$2
`

type AuthCodeUserParams struct {
	AuthCode sql.NullString `json:"auth_code"`
	ID       int32          `json:"id"`
}

func (q *Queries) AuthCodeUser(ctx context.Context, arg AuthCodeUserParams) error {
	_, err := q.db.ExecContext(ctx, authCodeUser, arg.AuthCode, arg.ID)
	return err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users as u(username,password,full_name,auth_code,email) 
VALUES ($1, $2, $3, $4,$5) 
RETURNING id, username, password, auth_code, full_name, email, password_changed_date, updated_date, created_date
`

type CreateUserParams struct {
	Username string         `json:"username"`
	Password string         `json:"password"`
	FullName string         `json:"full_name"`
	AuthCode sql.NullString `json:"auth_code"`
	Email    string         `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Password,
		arg.FullName,
		arg.AuthCode,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.AuthCode,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedDate,
		&i.UpdatedDate,
		&i.CreatedDate,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, password, auth_code, full_name, email, password_changed_date, updated_date, created_date from users WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.AuthCode,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedDate,
		&i.UpdatedDate,
		&i.CreatedDate,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
select id, username, password, auth_code, full_name, email, password_changed_date, updated_date, created_date from users
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Password,
			&i.AuthCode,
			&i.FullName,
			&i.Email,
			&i.PasswordChangedDate,
			&i.UpdatedDate,
			&i.CreatedDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}