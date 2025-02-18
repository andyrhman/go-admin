package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FullName string    `json:"fullName"`
	Username string    `json:"username"`
	Email    string    `json:"email" gorm:"unique"`
	Password []byte    `json:"-"`
}
