package services

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/lin-sel/bookmark-app/models"
	"github.com/lin-sel/bookmark-app/repository"
	"github.com/lin-sel/bookmark-app/web"
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
	// authuser := models.User{}
	// err := auth.Repository.GetByField(uow, user.Getusername(), "username", &authuser, []string{})
	// fmt.Println(authuser)
	// if !authuser.IsEmpty() {
	// 	return errors.New("user Already present")
	// }
	err := auth.Repository.Add(uow, user)
	if err != nil {
		uow.Complete()
		if strings.Contains(err.Error(), "Duplicate entry") {
			return web.NewValidationError("user", map[string]string{"error": "user already present"})
		}
		return err
	}
	uow.Commit()
	return err
}

// Login Return Auth User.
func (auth *UserService) Login(user *models.User) error {
	uow := repository.NewUnitOfWork(auth.DB, true)
	authuser := models.User{}
	err := auth.Repository.GetByField(uow, user.Getusername(), "username", &authuser, []string{"Category"})
	if err := authuser.IsUserValid(); err != nil {
		return errors.New("Invalid Username")
	}
	if err != nil {
		uow.Complete()
		return err
	}

	if authuser.GetAttemptime() >= 3 {
		return errors.New("You attemp to login account number of time with wrong password now your account has blocked. contact Admin")
	}
	if authuser.Getpassword() != user.Getpassword() {
		authuser.Attemptime = authuser.GetAttemptime() + 1
		auth.Update(&authuser)
		return errors.New("Invalid password")
	}

	authuser.Attemptime = 0
	auth.Update(&authuser)
	uow.Commit()
	*user = authuser
	return err
}

// Get Return User.
func (auth *UserService) Get(uid *uuid.UUID, user *[]models.User) error {
	uow := repository.NewUnitOfWork(auth.DB, true)
	err := auth.Repository.Get(uow, user, *uid, []string{})
	if err != nil {
		uow.Complete()
		return err
	}
	// uow.Commit()
	return err
}

// GetAll Return All User.
func (auth *UserService) GetAll(uid *uuid.UUID, user *[]models.User) error {
	uow := repository.NewUnitOfWork(auth.DB, true)
	err := auth.Repository.GetAll(uow, *uid, user, []string{})
	if err != nil {
		uow.Complete()
		return err
	}
	// uow.Commit()
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
	// auth.Repository.Delete(uow, )
	err := auth.Repository.Delete(uow, *uid, &models.User{})
	if err != nil {
		uow.Complete()
		return err
	}
	uow.Commit()
	return err
}

// CheckUser Parse UserID and check there model
func (auth *UserService) CheckUser(param map[string]string, role string) (*uuid.UUID, error) {
	userid, err := web.ParseID(param["userid"])
	if err != nil {
		return nil, web.NewValidationError("User ID", map[string]string{"error": "Invalid Admin ID"})
	}
	users := []models.User{*models.NewUserWithID(*userid)}
	err = auth.Get(userid, &users)
	if err != nil {
		return nil, err
	}
	if len(users) <= 0 {
		return nil, web.NewValidationError("User ID", map[string]string{"error": "Invalid Admin ID"})
	}
	if err := users[0].IsUserValid(); err != nil {
		return nil, web.NewValidationError("User ID", map[string]string{"error": "Invalid Admin ID"})
	}

	if !users[0].IsEqualRole(role) {
		return nil, web.NewValidationError("User ID", map[string]string{"error": "Invalid user"})
	}
	return userid, nil
}
