package auth

import (
	"github.com/google/uuid"
	"github.com/lin-sel/bookmark-app/basemodel"
	"github.com/lin-sel/bookmark-app/bookmark"
)

// User define Name, Username, Password
type User struct {
	basemodel.Basemodel
	Name     string               `gorm:"type:varchar(30);"`
	Username string               `gorm:"type:varchar(25);unique_index"`
	Password string               `gorm:"type:varchar(70)"`
	Bookmark *[]bookmark.Bookmark `json:"-"`
}

// NewUser Return New Object of User
func NewUser(name string, username string, password string) *User {
	return &User{
		Basemodel: basemodel.Basemodel{
			ID: GetUUID(),
		},
		Bookmark: nil,
		Name:     name,
		Password: password,
		Username: username,
	}
}

// Getusername return username
func (user *User) Getusername() string {
	return user.Username
}

// Getpassword return password
func (user *User) Getpassword() string {
	return user.Password
}

// Getname return name
func (user *User) Getname() string {
	return user.Name
}

// IsEmpty Return true/False
func (user *User) IsEmpty() bool {
	if user.Getusername() == "" {
		return true
	}
	return false
}

// GetUUID return uuid
func GetUUID() uuid.UUID {
	id, err := uuid.NewRandom()
	if err != nil {
		return GetUUID()
	}
	return id
}
