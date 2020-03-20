package services

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/lin-sel/bookmark-app/models"
	"github.com/lin-sel/bookmark-app/repository"
)

// UserService Structure
type UserService struct {
	DB         *gorm.DB
	Repository *repository.Repositorysrv
}

// NewUserService Return AuthService
func NewUserService(db *gorm.DB, repo *repository.Repositorysrv) *UserService {
	db.AutoMigrate(models.User{})
	return &UserService{
		DB:         db,
		Repository: repo,
	}
}

// Register User
func (auth *UserService) Register(user *models.User) error {
	uow := repository.NewUnitOfWork(auth.DB, false)
	authuser := models.User{}
	err := auth.Repository.GetByField(uow, user.Getusername(), "username", "", &authuser, []string{})
	if !authuser.IsEmpty() {
		return errors.New("user Already present")
	}
	err = auth.Repository.Add(uow, user)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// Login Return Auth User.
func (auth *UserService) Login(user *models.User) error {
	uow := repository.NewUnitOfWork(auth.DB, true)
	authuser := models.User{}
	err := auth.Repository.GetByField(uow, user.Getusername(), "username", "", &authuser, []string{"Category"})
	if authuser.IsEmpty() {
		return errors.New("Invalid User")
	}
	if err != nil {
		uow.Complete()
		return err
	}
	if !checkUserCreadential(user, &authuser) {
		return errors.New("Invalid User")
	}
	// uow.Commit()
	*user = authuser
	return err
}

func checkUserCreadential(user *models.User, authuser *models.User) bool {
	if user.Getpassword() == authuser.Getpassword() {
		return true
	}
	return false
}
