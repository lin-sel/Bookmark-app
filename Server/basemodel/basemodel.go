package basemodel

import (
	"time"

	uuid "github.com/google/uuid"
)

// Basemodel Define ID, createdAt, deletedAt, DeletedAt.
type Basemodel struct {
	ID        uuid.UUID  `gorm:"type:varchar(36);primary_key;"`
	CreatedAt time.Time  `gorm:"column:createdOn" json:"-"`
	UpdatedAt time.Time  `gorm:"column:updatedOn" json:"-"`
	DeletedAt *time.Time `sql:"index" gorm:"column:deletedOn" json:"-"`
}
