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
	GetTotalCount(ufw *UnitOfWork, id uuid.UUID, out interface{}) error
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
	return errors.New("Bad Request")
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
	case *models.BookmarkResponse:
		response := out.(*models.BookmarkResponse)
		return db.Debug().Model(response.ListOfBookmarkDTO).Limit(response.PageSize).Offset((response.PageNumber*response.PageSize)-response.PageSize).Find(response.ListOfBookmarkDTO, "user_id = ?", (*response.ListOfBookmarkDTO)[0].UserID).Error
	case *models.CategoryResponse:
		response := out.(*models.CategoryResponse)
		return db.Debug().Model(response.ListOfCategoryDTO).Limit(response.PageSize).Offset((response.PageNumber*response.PageSize)-response.PageSize).Find(response.ListOfCategoryDTO, "user_id = ?", (*response.ListOfCategoryDTO)[0].UserID).Error
		// return db.Model(out).Debug().Limit(2).Offset(0).Find(out, "user_id = ?", (*category)[0].UserID).Error
	}
	return errors.New("Bad Request")
}

// Update Entity
func (repo *Repositorysrv) Update(ufw *UnitOfWork, entity interface{}) error {
	return ufw.DB.Model(entity).Save(entity).Error
}

// Delete Entity From Database
func (repo *Repositorysrv) Delete(ufw *UnitOfWork, id uuid.UUID, out interface{}) error {
	switch out.(type) {
	case *models.Bookmark:
		bookmark := out.(*models.Bookmark)
		return ufw.DB.Debug().Model(out).Delete(out, "id = ? and user_id = ?", id, (*bookmark).UserID).Error
	case *models.User:
		return ufw.DB.Debug().Model(out).Delete(out, "id = ?", id).Error
	case *models.Category:
		category := out.(*models.Category)
		bookmark := models.Bookmark{}
		ufw.DB.Model(bookmark).Delete(bookmark, "category_id = ? and user_id = ?", id, (*category).UserID)
		return ufw.DB.Debug().Model(out).Delete(out, "id = ? and user_id = ?", id, (*category).UserID).Error
	}
	return errors.New("Bad Request")
}

// DeleteByField Delete Based On Column Condition
func (repo *Repositorysrv) DeleteByField(ufw *UnitOfWork, value interface{}, fieldname string, id uuid.UUID, out interface{}) {
	ufw.DB.Model(out).Delete(out, fmt.Sprintf("%s = ? and user_id = ?", fieldname), id)
}

// GetByField Return Result Based on Field.
func (repo *Repositorysrv) GetByField(ufw *UnitOfWork, value interface{}, fieldname string, out interface{}, preloadAssociation []string) error {
	db := ufw.DB
	for _, association := range preloadAssociation {
		db = db.Preload(association)
	}
	switch out.(type) {
	case *models.User:
		return db.Model(out).Debug().First(out, fmt.Sprintf("%s = ?", fieldname), value).Error
	case *models.BookmarkResponse:
		response := out.(*models.BookmarkResponse)
		return db.Debug().Model(response.ListOfBookmarkDTO).Limit(response.PageSize).Offset((response.PageNumber*response.PageSize)-response.PageSize).Find(response.ListOfBookmarkDTO, fmt.Sprintf("%s = ? and user_id = ?", fieldname), value, (*response.ListOfBookmarkDTO)[0].UserID).Error
	case *[]models.Category:
		category := out.(*[]models.Category)
		return db.Model(out).Debug().First(out, fmt.Sprintf("%s = ? and user_id = ?", fieldname), value, (*category)[0].UserID).Error
	}
	return errors.New("Bad Request")
}

// GetTotalCount return total dataset.
func (repo *Repositorysrv) GetTotalCount(ufw *UnitOfWork, id uuid.UUID, out interface{}, count interface{}, fieldname string, fieldvalue interface{}) error {
	switch out.(type) {
	case *models.Bookmark:
		bookmark := out.(*models.Bookmark)
		return ufw.DB.Debug().Model(out).Where(fmt.Sprintf("%s = ? and user_id = ?", fieldname), fieldvalue, bookmark.GetUserID()).Count(count).Error
	case *models.Category:
		category := out.(*models.Category)
		return ufw.DB.Debug().Model(out).Where("user_id = ?", category.GetUserID()).Count(count).Error
	case *models.User:
		return ufw.DB.Debug().Model(out).Count(count).Error
	}
	return nil
}
