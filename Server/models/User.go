package models

import (
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
	Attemptime int8        `gorm:"type:int" json:"-"`
	Category   []Category  `json:"categories"`
}

// NewUserWithID Return user instance with ID
func NewUserWithID() *User {
	return &User{
		Basemodel: Basemodel{
			ID: web.GetUUID(),
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

// GetAttemptime Return number of time attemp.
func (user *User) GetAttemptime() int8 {
	return user.Attemptime
}

// IsEmpty Return true/False
func (user *User) IsEmpty() bool {
	if user.Getusername() == "" {
		return true
	}
	return false
}
