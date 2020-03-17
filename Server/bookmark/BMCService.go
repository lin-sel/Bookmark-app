package bookmark

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lin-sel/bookmark-app/repository"
)

//BMCService Bookmark Category Service
type BMCService struct {
	DB         *gorm.DB
	Repository *repository.Repositorysrv
}

// NewBMCService Return New Service Object
func NewBMCService(repo *repository.Repositorysrv, db *gorm.DB) *BMCService {
	db.AutoMigrate(Category{})
	return &BMCService{Repository: repo, DB: db}
}

// AddBMCategory to Database
func (bm *BMCService) AddBMCategory(category *Category) error {
	uow := repository.NewUnitOfWork(bm.DB, false)
	err := bm.Repository.Add(uow, category)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// RecentBMCategory to Database
func (bm *BMCService) RecentBMCategory() {

}

// UpdateBMCategory to Database
func (bm *BMCService) UpdateBMCategory(category *Category) error {
	uow := repository.NewUnitOfWork(bm.DB, false)
	err := bm.Repository.Update(uow, category)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// DeleteBMCategory to Database
func (bm *BMCService) DeleteBMCategory(uid, cid uuid.UUID) error {
	uow := repository.NewUnitOfWork(bm.DB, true)
	err := bm.Repository.Delete(uow, uid, cid, Category{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// GetBMCategory to Database
func (bm *BMCService) GetBMCategory(uid, cid uuid.UUID, category *[]Category) error {
	uow := repository.NewUnitOfWork(bm.DB, true)
	err := bm.Repository.Get(uow, category, uid, cid, []string{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// GetAllBMCategory From database
func (bm *BMCService) GetAllBMCategory(uid uuid.UUID, categories *[]Category) error {
	uow := repository.NewUnitOfWork(bm.DB, true)
	err := bm.Repository.GetAll(uow, uid, categories, []string{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}
