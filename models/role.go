package models

type Role struct {
	Id          uint         `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permission" gorm:"many2many:role_permissions;constraint:OnDelete:CASCADE;"`
}
