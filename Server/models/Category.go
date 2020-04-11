package models

import (
	uuid "github.com/satori/go.uuid"
)

// Category Structure
type Category struct {
	Basemodel
	CName    string     `gorm:"type:varchar(100);unique_index:name_userid" json:"category"`
	UserID   uuid.UUID  `gorm:"type:varchar(40);not_null;unique_index:name_userid" json:"-"`
	Bookmark []Bookmark `json:"-"`
}

// NewCategoryWithUserID Return Category Object
func NewCategoryWithUserID(userid uuid.UUID) *Category {
	return &Category{
		UserID: userid,
	}
}

//GetCategoryName return name
func (category *Category) GetCategoryName() string {
	return category.CName
}

//GetCategoryID return ID of Category
func (category *Category) GetCategoryID() uuid.UUID {
	return category.ID
}

//GetUserID return ID of User to which category Belong
func (category *Category) GetUserID() uuid.UUID {
	return category.UserID
}
