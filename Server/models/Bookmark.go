package models

import (
	"github.com/lin-sel/bookmark-app/web"
	uuid "github.com/satori/go.uuid"
)

// Bookmark Structure
type Bookmark struct {
	Basemodel
	Label      string    `gorm:"type:varchar(200)" json:"label"`
	URL        string    `gorm:"type:varchar(300)" json:"url"`
	Tag        string    `gorm:"type:varchar(300)" json:"tag"`
	UserID     uuid.UUID `gorm:"type:varchar(40);not_null" json:"-"`
	CategoryID uuid.UUID `gorm:"type:varchar(40);not_null" json:"categoryid"`
}

// NewBookmark return Bookmark Struct
func NewBookmark(label string, url string, tag string, userid uuid.UUID) *Bookmark {
	return &Bookmark{
		Basemodel: Basemodel{
			ID: web.GetUUID(),
		},
		Label:  label,
		Tag:    tag,
		URL:    url,
		UserID: userid,
	}
}

// NewBookmarkWithID return New Bookmark Instance with ID
func NewBookmarkWithID() *Bookmark {
	return &Bookmark{
		Basemodel: Basemodel{
			ID: web.GetUUID(),
		},
	}
}

// GetLabel return Label of Bookmark
func (bookmark *Bookmark) GetLabel() string {
	return bookmark.Label
}

// GetURL return URL
func (bookmark *Bookmark) GetURL() string {
	return bookmark.URL
}

// GetTag Return Tag of Bookmark
func (bookmark *Bookmark) GetTag() string {
	return bookmark.Tag
}

//GetUserID return id
func (bookmark *Bookmark) GetUserID() uuid.UUID {
	return bookmark.UserID
}

//GetCategoryID Return ID
func (bookmark *Bookmark) GetCategoryID() uuid.UUID {
	return bookmark.CategoryID
}
