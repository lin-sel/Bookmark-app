package repository

import (
	"errors"
	"fmt"

	"github.com/lin-sel/bookmark-app/models"
	uuid "github.com/satori/go.uuid"
)

// Repository Struct
type Repository interface {
	GetAll(ufw *UnitOfWork, id uuid.UUID, out interface{}, preloadAssociation []string) error
	Get(ufw *UnitOfWork, out interface{}, id uuid.UUID, preloadAssociation []string) error
	Add(ufw *UnitOfWork, input interface{}) error
	Delete(ufw *UnitOfWork, id uuid.UUID, out interface{}) error
	Update(ufw *UnitOfWork, id uuid.UUID, out interface{}) error
	GetByField(ufw *UnitOfWork, value interface{}, fieldname string, preloadAssociation []string) error
}

// Repositorysrv Return new Service
type Repositorysrv struct {
}

// NewRepository Return New Object
func NewRepository() *Repositorysrv {
	return &Repositorysrv{}
}

// Add New Student To Database
func (repo *Repositorysrv) Add(ufw *UnitOfWork, out interface{}) error {
	return ufw.DB.Create(out).Error
}

// Get Entity By ID
func (repo *Repositorysrv) Get(ufw *UnitOfWork, out interface{}, id uuid.UUID, preloadAssociation []string) error {
	db := ufw.DB
	for _, association := range preloadAssociation {
		db = db.Preload(association)
	}

	switch out.(type) {
	case *[]models.User:
		return ufw.DB.Model(out).Debug().First(out, "id = ?", id).Error
	case *[]models.Bookmark:
		bookmark := out.(*[]models.Bookmark)
		return db.Model(out).Debug().First(out, "id = ? and user_id = ?", id, (*bookmark)[0].UserID).Error
	case *[]models.Category:
		category := out.(*[]models.Category)
		return db.Model(out).Debug().First(out, "id = ? and user_id = ?", id, (*category)[0].UserID).Error
	}
	return errors.New("Unknown Error Occur")
}

// GetAll Entity
func (repo *Repositorysrv) GetAll(ufw *UnitOfWork, id uuid.UUID, out interface{}, preloadAssociation []string) error {
	db := ufw.DB
	for _, association := range preloadAssociation {
		db = db.Preload(association)
	}
	switch out.(type) {
	case *[]models.User:
		return db.Model(out).Debug().Find(out).Error
	case *[]models.Bookmark:
		bookmark := out.(*[]models.Bookmark)
		fmt.Println((*bookmark)[0].UserID)
		return db.Model(out).Debug().Find(out, "user_id = ?", (*bookmark)[0].UserID).Error
	case *[]models.Category:
		category := out.(*[]models.Category)
		return db.Model(out).Debug().Find(out, "user_id = ?", (*category)[0].UserID).Error
	}
	return errors.New("Unknown Error Occur")
}

// Update Entity
func (repo *Repositorysrv) Update(ufw *UnitOfWork, entity interface{}) error {
	return ufw.DB.Model(entity).Save(entity).Error
}

// Delete Entity From Database
func (repo *Repositorysrv) Delete(ufw *UnitOfWork, id uuid.UUID, out interface{}) error {
	switch out.(type) {
	case *models.User:
		return ufw.DB.Debug().Model(out).Delete(out, "id = ?", id).Error
	case *models.Bookmark:
		bookmark := out.(*models.Bookmark)
		return ufw.DB.Debug().Model(out).Delete(out, "id = ? and user_id = ?", id, (*bookmark).UserID).Error
	case *models.Category:
		category := out.(*models.Category)
		return ufw.DB.Debug().Model(out).Delete(out, "id = ? and user_id = ?", id, (*category).UserID).Error
	}
	return errors.New("Unknown error Occur")
}

// GetByField Return Result Based on Field.
func (repo *Repositorysrv) GetByField(ufw *UnitOfWork, value interface{}, fieldname string, out interface{}, preloadAssociation []string) error {
	db := ufw.DB
	for _, association := range preloadAssociation {
		db = db.Preload(association)
	}
	switch out.(type) {
	case *models.User, *models.Admin:
		return db.Model(out).Debug().First(out, fmt.Sprintf("%s = ?", fieldname), value).Error
	case *[]models.Bookmark:
		bookmark := out.(*[]models.Bookmark)
		return db.Model(out).Debug().Find(out, fmt.Sprintf("%s = ? and user_id = ?", fieldname), value, (*bookmark)[0].UserID).Error
	case *[]models.Category:
		category := out.(*[]models.Category)
		return db.Model(out).Debug().First(out, fmt.Sprintf("%s = ? and user_id = ?", fieldname), value, (*category)[0].UserID).Error
	}
	return errors.New("Unknown error Occur")
}
