// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"database/sql"
	"time"
)

type User struct {
	ID                  int32          `json:"id"`
	Username            string         `json:"username"`
	Password            string         `json:"password"`
	AuthCode            sql.NullString `json:"auth_code"`
	FullName            string         `json:"full_name"`
	Email               string         `json:"email"`
	PasswordChangedDate time.Time      `json:"password_changed_date"`
	UpdatedDate         sql.NullTime   `json:"updated_date"`
	CreatedDate         time.Time      `json:"created_date"`
}
