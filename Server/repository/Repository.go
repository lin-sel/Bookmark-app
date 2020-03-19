package repository

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// Repository Struct
type Repository interface {
	GetAll(ufw *UnitOfWork, uid uuid.UUID, out interface{}, preloadAssociation []string) error
	Get(ufw *UnitOfWork, out interface{}, uid uuid.UUID, bid uuid.UUID, preloadAssociation []string) error
	Add(ufw *UnitOfWork, input interface{}) error
	Delete(ufw *UnitOfWork, uid uuid.UUID, bid uuid.UUID, out interface{}) error
	Update(ufw *UnitOfWork, uid uuid.UUID, out interface{}) error
	GetByField(ufw *UnitOfWork, value interface{}, condition string, preloadAssociation []string) error
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

// Get Student By ID
func (repo *Repositorysrv) Get(ufw *UnitOfWork, out interface{}, uid, bid uuid.UUID, preloadAssociation []string) error {
	db := ufw.DB
	for _, association := range preloadAssociation {
		db = db.Preload(association)
	}
	return db.Debug().Model(out).First(out, "id = ? and user_id = ?", bid, uid).Error
}

// GetAll Student
func (repo *Repositorysrv) GetAll(ufw *UnitOfWork, uid uuid.UUID, out interface{}, preloadAssociation []string) error {
	db := ufw.DB
	for _, association := range preloadAssociation {
		db = db.Preload(association)
	}
	return db.Model(out).Debug().Find(out, "user_id = ?", uid).Error
}

// Update Student
func (repo *Repositorysrv) Update(ufw *UnitOfWork, entity interface{}) error {
	return ufw.DB.Model(entity).Save(entity).Error
}

// Delete Student From Database
func (repo *Repositorysrv) Delete(ufw *UnitOfWork, uid, bid uuid.UUID, out interface{}) error {
	return ufw.DB.Debug().Model(out).Delete(out, "user_id = ? and id = ?", uid, bid).Error
}

// GetByField Return Result Based on Field.
func (repo *Repositorysrv) GetByField(ufw *UnitOfWork, value interface{}, fieldname string, uid interface{}, out interface{}, preloadAssociation []string) error {
	db := ufw.DB
	for _, association := range preloadAssociation {
		db = db.Preload(association)
	}
	if uid == "" {
		return db.Model(out).Debug().First(out, fmt.Sprintf("%s = ?", fieldname), value).Error
	}
	return db.Model(out).First(out, fmt.Sprintf("%s = ? and user_id = ?", fieldname), value, uid).Error
}
