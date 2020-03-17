package auth

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/lin-sel/bookmark-app/repository"
)

// Authsrv Structure
type Authsrv struct {
	DB         *gorm.DB
	Repository *repository.Repositorysrv
}

// NewAuthsrv Return AuthService
func NewAuthsrv(db *gorm.DB, repo *repository.Repositorysrv) *Authsrv {
	db.AutoMigrate(User{})
	return &Authsrv{
		DB:         db,
		Repository: repo,
	}
}

// Register User
func (auth *Authsrv) Register(user *User) error {
	uow := repository.NewUnitOfWork(auth.DB, true)
	authuser := User{}
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
func (auth *Authsrv) Login(user *User) error {
	uow := repository.NewUnitOfWork(auth.DB, true)
	authuser := User{}
	err := auth.Repository.GetByField(uow, user.Getusername(), "username", "", &authuser, []string{})
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
	uow.Commit()
	return err
}

func checkUserCreadential(user *User, authuser *User) bool {
	if user.Getpassword() == authuser.Getpassword() {
		return true
	}
	return false
}
