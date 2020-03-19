package services

import (
	"github.com/jinzhu/gorm"
	"github.com/lin-sel/bookmark-app/models"
	"github.com/lin-sel/bookmark-app/repository"
	uuid "github.com/satori/go.uuid"
)

// BookmarkService Struct
type BookmarkService struct {
	DB         *gorm.DB
	Repository *repository.Repositorysrv
}

// NewBookmarkService return
func NewBookmarkService(repo *repository.Repositorysrv, db *gorm.DB) *BookmarkService {
	db.AutoMigrate(models.Bookmark{})
	db.Model(models.Bookmark{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(models.Bookmark{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
	return &BookmarkService{Repository: repo, DB: db}
}

// AddBookmark to Database
func (bm *BookmarkService) AddBookmark(bookmark *models.Bookmark) error {
	uow := repository.NewUnitOfWork(bm.DB, false)
	err := bm.Repository.Add(uow, bookmark)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// RecentBookmark to Database
func (bm *BookmarkService) RecentBookmark() {

}

// UpdateBookmark to Database
func (bm *BookmarkService) UpdateBookmark(bookmark *models.Bookmark) error {
	uow := repository.NewUnitOfWork(bm.DB, false)
	err := bm.Repository.Update(uow, bookmark)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// DeleteBookmark to Database
func (bm *BookmarkService) DeleteBookmark(uid, bid uuid.UUID) error {
	uow := repository.NewUnitOfWork(bm.DB, false)
	err := bm.Repository.Delete(uow, uid, bid, models.Bookmark{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// GetBookmark to Database
func (bm *BookmarkService) GetBookmark(uid, bid uuid.UUID, bookmark *[]models.Bookmark) error {
	uow := repository.NewUnitOfWork(bm.DB, true)
	err := bm.Repository.Get(uow, bookmark, uid, bid, []string{})
	if err != nil {
		uow.Complete()
		return err
	}
	// uow.Commit()
	return err
}

// GetAllBookmark From database
func (bm *BookmarkService) GetAllBookmark(uid uuid.UUID, bookmark *[]models.Bookmark) error {
	uow := repository.NewUnitOfWork(bm.DB, true)
	err := bm.Repository.GetAll(uow, uid, bookmark, []string{})
	if err != nil {
		uow.Complete()
		return err
	}
	// uow.Commit()
	return err
}

// GetBookmarkByCategory From database
func (bm *BookmarkService) GetBookmarkByCategory(uid, cid uuid.UUID, bookmark *[]models.Bookmark) error {
	uow := repository.NewUnitOfWork(bm.DB, true)
	err := bm.Repository.GetByField(uow, cid, "category_id", uid, bookmark, []string{})
	if err != nil {
		uow.Complete()
		return err
	}
	// uow.Commit()
	return err
}
