package bookmark

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lin-sel/bookmark-app/repository"
)

// BMService Struct
type BMService struct {
	DB         *gorm.DB
	Repository *repository.Repositorysrv
}

// NewBMService return
func NewBMService(repo *repository.Repositorysrv, db *gorm.DB) *BMService {
	db.AutoMigrate(Bookmark{})
	db.Model(Bookmark{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(Bookmark{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
	return &BMService{Repository: repo, DB: db}
}

// AddBookmark to Database
func (bm *BMService) AddBookmark(bookmark *Bookmark) error {
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
func (bm *BMService) RecentBookmark() {

}

// UpdateBookmark to Database
func (bm *BMService) UpdateBookmark(bookmark *Bookmark) error {
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
func (bm *BMService) DeleteBookmark(uid, bid uuid.UUID) error {
	uow := repository.NewUnitOfWork(bm.DB, true)
	err := bm.Repository.Delete(uow, uid, bid, Bookmark{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// GetBookmark to Database
func (bm *BMService) GetBookmark(uid, bid uuid.UUID, bookmark *[]Bookmark) error {
	uow := repository.NewUnitOfWork(bm.DB, true)
	err := bm.Repository.Get(uow, bookmark, uid, bid, []string{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// GetAllBookmark From database
func (bm *BMService) GetAllBookmark(uid uuid.UUID, bookmark *[]Bookmark) error {
	uow := repository.NewUnitOfWork(bm.DB, true)
	err := bm.Repository.GetAll(uow, uid, bookmark, []string{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// GetBookmarkByCategory From database
func (bm *BMService) GetBookmarkByCategory(uid, cid uuid.UUID, bookmark *[]Bookmark) error {
	uow := repository.NewUnitOfWork(bm.DB, true)
	err := bm.Repository.GetByField(uow, cid, "category_id", uid, bookmark, []string{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}
