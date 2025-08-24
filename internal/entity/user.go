package entity

import (
	"time"
)

type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Role         string    `json:"role"`
	FullName     string    `json:"full_name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserRole string

const (
	RoleWriter UserRole = "writer"
	RoleEditor UserRole = "editor"
	RoleAdmin  UserRole = "admin"
)

func (User) TableName() string {
	return "users"
}
