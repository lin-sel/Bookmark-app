package services

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/lin-sel/bookmark-app/models"
	"github.com/lin-sel/bookmark-app/repository"
	uuid "github.com/satori/go.uuid"
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
		return errors.New("Invalid Username")
	}
	if err != nil {
		uow.Complete()
		return err
	}

	if authuser.GetAttemptime() >= 3 {
		return errors.New("You attemp to login account number of time with wrong password now your account has blocked. contact Admin")
	}
	if !checkUserCreadential(user, &authuser) {
		authuser.Attemptime = authuser.GetAttemptime() + 1
		auth.Update(&authuser)
		return errors.New("Invalid Credential to Login")
	}

	authuser.Attemptime = 0
	auth.Update(&authuser)
	uow.Commit()
	*user = authuser
	return err
}

// Get Return User.
func (auth *UserService) Get(uid *uuid.UUID, user *models.User) error {
	uow := repository.NewUnitOfWork(auth.DB, true)
	err := auth.Repository.Get(uow, user, uid, "", []string{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// Update Return Auth User.
func (auth *UserService) Update(user *models.User) error {
	uow := repository.NewUnitOfWork(auth.DB, true)
	err := auth.Repository.Update(uow, user)
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// Delete Return Auth User.
func (auth *UserService) Delete(uid *uuid.UUID) error {
	uow := repository.NewUnitOfWork(auth.DB, true)
	err := auth.Repository.Delete(uow, uid, "", models.User{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

func checkUserCreadential(user *models.User, authuser *models.User) bool {
	if user.Getpassword() == authuser.Getpassword() {
		return true
	}
	return false
}
