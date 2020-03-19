package services

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lin-sel/bookmark-app/models"
	"github.com/lin-sel/bookmark-app/repository"
)

//BookmarkCategoryService Bookmark Category Service
type BookmarkCategoryService struct {
	DB         *gorm.DB
	Repository *repository.Repositorysrv
}

// NewBookmarkCategoryService Return New Service Object
func NewBookmarkCategoryService(repo *repository.Repositorysrv, db *gorm.DB) *BookmarkCategoryService {
	db.AutoMigrate(models.Category{})
	return &BookmarkCategoryService{Repository: repo, DB: db}
}

// AddBookmarkCategory to Database
func (bm *BookmarkCategoryService) AddBookmarkCategory(category *models.Category) error {
	uow := repository.NewUnitOfWork(bm.DB, false)
	err := bm.Repository.Add(uow, category)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// RecentBookmarkCategory to Database
func (bm *BookmarkCategoryService) RecentBookmarkCategory() {

}

// UpdateBookmarkCategory to Database
func (bm *BookmarkCategoryService) UpdateBookmarkCategory(category *models.Category) error {
	uow := repository.NewUnitOfWork(bm.DB, false)
	err := bm.Repository.Update(uow, category)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// DeleteBookmarkCategory to Database
func (bm *BookmarkCategoryService) DeleteBookmarkCategory(uid, cid uuid.UUID) error {
	uow := repository.NewUnitOfWork(bm.DB, true)
	err := bm.Repository.Delete(uow, uid, cid, models.Category{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// GetBookmarkCategory to Database
func (bm *BookmarkCategoryService) GetBookmarkCategory(uid, cid uuid.UUID, category *[]models.Category) error {
	uow := repository.NewUnitOfWork(bm.DB, true)
	err := bm.Repository.Get(uow, category, uid, cid, []string{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// GetAllBookmarkCategory From database
func (bm *BookmarkCategoryService) GetAllBookmarkCategory(uid uuid.UUID, categories *[]models.Category) error {
	uow := repository.NewUnitOfWork(bm.DB, true)
	err := bm.Repository.GetAll(uow, uid, categories, []string{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}
