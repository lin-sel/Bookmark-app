package services

import (
	"github.com/jinzhu/gorm"
	"github.com/lin-sel/bookmark-app/models"
	"github.com/lin-sel/bookmark-app/repository"
	"github.com/lin-sel/bookmark-app/web"
)

// AdminService Struct Get Control On All Function
type AdminService struct {
	DB         *gorm.DB
	Repository *repository.Repositorysrv
}

// NewAdminService Return New Instance Of AdminService
func NewAdminService(db *gorm.DB, repo *repository.Repositorysrv) *AdminService {
	db.AutoMigrate(models.Admin{})
	return &AdminService{
		DB:         db,
		Repository: repo,
	}
}

//Login admin Login
func (service *AdminService) Login(requestedadmin *models.Admin) error {
	ufw := repository.NewUnitOfWork(service.DB, true)
	var authadmin models.Admin
	err := service.Repository.GetByField(ufw, requestedadmin.Getusername(), "username", &authadmin, []string{})
	if err != nil {
		return web.NewValidationError("username", map[string]string{"error": "Invalid username"})
	}
	if requestedadmin.Getpassword() != authadmin.Getpassword() {
		return web.NewValidationError("password", map[string]string{"error": "Invalid password"})

	}
	*requestedadmin = authadmin
	return nil
}
