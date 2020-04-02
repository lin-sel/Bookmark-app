package models

import (
	"strings"

	"github.com/lin-sel/bookmark-app/web"
	uuid "github.com/satori/go.uuid"
)

// User define Name, Username, Password
type User struct {
	Basemodel
	Name       string      `gorm:"type:varchar(30);" json:"name"`
	Username   string      `gorm:"type:varchar(25);unique_index" json:"username"`
	Password   string      `gorm:"type:varchar(70)" json:"password"`
	Email      string      `gorm:"type:varchar(30)" json:"email"`
	Profile    interface{} `gorm:"type:LONGBLOB" json:"profile"`
	Attemptime int8        `gorm:"type:int" json:"attempt"`
	Category   []Category  `json:"categories"`
}

// NewUserWithNewID Return user instance with ID
func NewUserWithNewID() *User {
	return &User{
		Basemodel: Basemodel{
			ID: web.GetUUID(),
		},
	}
}

// NewUserWithID Return user with existing user id.
func NewUserWithID(uid uuid.UUID) *User {
	return &User{
		Basemodel: Basemodel{
			ID: uid,
		},
	}
}

// NewUser Return New Object of User
func NewUser(name string, username string, password string) *User {
	return &User{
		Basemodel: Basemodel{
			ID: web.GetUUID(),
		},
		Name:     name,
		Password: password,
		Username: username,
	}
}

// Getusername return username
func (user *User) Getusername() string {
	return user.Username
}

// GetID return User ID
func (user *User) GetID() uuid.UUID {
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

// GetAttemptime Return number of time attemp.
func (user *User) GetAttemptime() int8 {
	return user.Attemptime
}

// IsEmpty Return true/False
func (user *User) IsEmpty() bool {
	if len(strings.Trim(user.Getusername(), " ")) == 0 || len(strings.Trim(user.Getpassword(), " ")) == 0 || len(strings.Trim(user.Getname(), " ")) == 0 {
		return true
	}
	return false
}

// IsUserValid check all nessessory paramater empty ? and return error
func (user *User) IsUserValid() error {
	if len(strings.Trim(user.Getusername(), " ")) == 0 {
		return web.NewValidationError("Require", map[string]string{"error": "username required"})
	}

	if len(strings.Trim(user.Getpassword(), " ")) == 0 {
		return web.NewValidationError("Require", map[string]string{"error": "password required"})
	}

	// if len(strings.Trim(user.Getname(), " ")) == 0 {
	// 	return web.NewValidationError("Require", map[string]string{"error": "name required"})
	// }
	return nil
}
