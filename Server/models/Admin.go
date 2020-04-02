package models

import (
	"github.com/lin-sel/bookmark-app/web"
	uuid "github.com/satori/go.uuid"
)

// IUser Interface Implement By User And Admin
type IUser interface {
	Getusername() string
	Getpassword() string
	GetID() uuid.UUID
}

// Admin Structure
type Admin struct {
	Basemodel
	Email    string `gorm:"varchar(40);unique_index" json:"email"`
	Username string `gorm:"varchar(40);unique_index" json:"username"`
	Password string `gorm:"varchar(40)" json:"password"`
}

// NewAdminWithNewID return New Admin Instance With New ID
func NewAdminWithNewID() *Admin {
	return &Admin{
		Basemodel: Basemodel{
			ID: web.GetUUID(),
		},
	}
}

// NewAdmin Return New Admin Instance With Proviced ID
func NewAdmin(id uuid.UUID) *Admin {
	return &Admin{
		Basemodel: Basemodel{
			ID: id,
		},
	}
}

// Getusername return Username
func (admin *Admin) Getusername() string {
	return admin.Username
}

// Getpassword return password
func (admin *Admin) Getpassword() string {
	return admin.Password
}

// GetEmail return Email
func (admin *Admin) GetEmail() string {
	return admin.Email
}

// GetID Return Admin ID
func (admin *Admin) GetID() uuid.UUID {
	return admin.ID
}
