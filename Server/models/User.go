package models

import (
	"github.com/google/uuid"
	"github.com/lin-sel/bookmark-app/web"
)

// User define Name, Username, Password
type User struct {
	Basemodel
	Name      string      `gorm:"type:varchar(30);"`
	Username  string      `gorm:"type:varchar(25);unique_index"`
	Password  string      `gorm:"type:varchar(70)"`
	Bookmarks *[]Bookmark `json:"-"`
}

// NewUser Return New Object of User
func NewUser(name string, username string, password string) *User {
	return &User{
		Basemodel: Basemodel{
			ID: web.GetUUID(),
		},
		Bookmarks: nil,
		Name:      name,
		Password:  password,
		Username:  username,
	}
}

// Getusername return username
func (user *User) Getusername() string {
	return user.Username
}

// GetuserID return User ID
func (user *User) GetuserID() uuid.UUID {
	return user.ID
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
