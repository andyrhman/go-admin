package models

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
	"gorm.io/gorm"
)

type User struct {
	Id       uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FullName string    `json:"fullName"`
	Username string    `json:"username"`
	Email    string    `json:"email" gorm:"unique"`
	Password []byte    `json:"-"`
	RoleId   uint      `json:"role_id"`
	Role     Role      `json:"role" gorm:"foreignKey:RoleId"`
}

func generateSalt() []byte {
	salt := make([]byte, 16)
	rand.Read(salt)
	return salt
}

func (user *User) SetPassword(password string) {
	salt := generateSalt()

	hashedPassword := argon2.IDKey([]byte(password), salt, 3, 64*1024, 4, 32)

	encodedSalt := base64.StdEncoding.EncodeToString(salt)
	encodedHash := base64.StdEncoding.EncodeToString(hashedPassword)

	user.Password = []byte(fmt.Sprintf("$argon2id$v=19$m=65536,t=3,p=4$%s$%s", encodedSalt, encodedHash))
}

func (user *User) ComparePassword(inputPassword string) bool {
	// Convert stored password (which is a []byte) to a string
	storedHash := string(user.Password)
	parts := strings.Split(storedHash, "$")
	if len(parts) != 6 {
		return false // Invalid hash format
	}

	// Extract salt & stored hash from the split parts
	encodedSalt := parts[4]
	encodedStoredHash := parts[5]

	// Decode salt from base64
	salt, err := base64.StdEncoding.DecodeString(encodedSalt)
	if err != nil {
		return false
	}

	// Hash the input password using the extracted salt
	newHash := argon2.IDKey([]byte(inputPassword), salt, 3, 64*1024, 4, 32)

	// Compare the computed hash with the stored hash
	return base64.StdEncoding.EncodeToString(newHash) == encodedStoredHash
}

func (users *User) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&User{}).Count(&total)

	return total
}

func (user *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	var users []User

	db.Preload("Role").Offset(offset).Limit(limit).Find(&users)

	return users
}
