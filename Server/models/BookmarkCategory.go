package models

import (
	"github.com/google/uuid"
	"github.com/lin-sel/bookmark-app/web"
)

// Category Structure
type Category struct {
	Basemodel
	CName  string    `gorm:"type:varchar(100)"`
	UserID uuid.UUID `gorm:"type:varchar(40);not_null"`
}

// NewCategory Return Category Object
func NewCategory(name string, userid uuid.UUID) *Category {
	return &Category{
		Basemodel: Basemodel{
			ID: web.GetUUID(),
		},
		CName:  name,
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