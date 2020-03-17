package bookmark

import (
	"github.com/google/uuid"
	"github.com/lin-sel/bookmark-app/basemodel"
)

// Bookmark Structure
type Bookmark struct {
	basemodel.Basemodel
	Label      string    `gorm:"type:varchar(200)"`
	URL        string    `gorm:"type:varchar(300)"`
	Tag        string    `gorm:"type:varchar(300)"`
	UserID     uuid.UUID `gorm:"type:varchar(40);not_null"`
	CategoryID uuid.UUID `gorm:"type:varchar(40);not_null"`
}

// NewBookmark return Bookmark Struct
func NewBookmark(label string, url string, tag string, userid uuid.UUID) *Bookmark {
	return &Bookmark{
		Basemodel: basemodel.Basemodel{
			ID: GetUUID(),
		},
		Label:  label,
		Tag:    tag,
		URL:    url,
		UserID: userid,
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

// GetUUID return uuid
func GetUUID() uuid.UUID {
	id, err := uuid.NewRandom()
	if err != nil {
		return GetUUID()
	}
	return id
}
