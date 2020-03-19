package repository

import "github.com/jinzhu/gorm"

// UnitOfWork Structure
type UnitOfWork struct {
	DB        *gorm.DB
	Readonly  bool
	Committed bool
}

// NewUnitOfWork Return New Object
func NewUnitOfWork(db *gorm.DB, readonly bool) *UnitOfWork {
	if readonly {
		return &UnitOfWork{DB: db.New(), Readonly: true, Committed: false}
	}
	return &UnitOfWork{DB: db.New().Begin(), Readonly: false, Committed: false}
}

// Complete For Rollback
func (uow *UnitOfWork) Complete() {
	if !uow.Readonly && !uow.Committed {
		uow.DB.Rollback()
	}
}

// Commit for Commit
func (uow *UnitOfWork) Commit() {
	if !uow.Readonly && !uow.Committed {
		uow.DB.Commit()
	}
}
