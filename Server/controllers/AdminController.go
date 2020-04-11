package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lin-sel/bookmark-app/models"
	"github.com/lin-sel/bookmark-app/services"
	"github.com/lin-sel/bookmark-app/web"
)

const adminRole = "admin"

// AdminController Handle All function of Admin Controller
type AdminController struct {
	bookmarkservice *services.BookmarkService
	categoryservice *services.BookmarkCategoryService
	userservice     *services.UserService
	adminservice    *services.AdminService
	authcontroller  *AuthController
}

// NewAdminController Return New Instance of AdminController.
func NewAdminController(bookmark *services.BookmarkService, category *services.BookmarkCategoryService, user *services.UserService, admin *services.AdminService, auth *AuthController) *AdminController {
	return &AdminController{
		bookmarkservice: bookmark,
		adminservice:    admin,
		authcontroller:  auth,
		categoryservice: category,
		userservice:     user,
	}
}

// AdminRouteRegister Register All Endpoint of Admin
func (admin *AdminController) AdminRouteRegister(r *mux.Router) {
	r.HandleFunc("/admin/login", admin.login).Methods("POST")
	s := r.PathPrefix("/admin/{userid}").Subrouter()
	s.Use(admin.authcontroller.AuthUser)
	s.HandleFunc("/user/{id}", admin.getUser).Methods("GET")
	s.HandleFunc("/user", admin.getAllUser).Methods("GET")
	s.HandleFunc("/user", admin.addNewUser).Methods("POST")
	s.HandleFunc("/user/{id}", admin.updateUser).Methods("PUT")
	s.HandleFunc("/user/{id}", admin.deleteUser).Methods("DELETE")
	s.HandleFunc("/user/{id}/bookmark", admin.allBookmark).Methods("GET")
	s.HandleFunc("/user/{id}/category", admin.allCategory).Methods("GET")
}

// login Admin Login
func (admin *AdminController) login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := web.UnmarshalJSON(r, &user)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("Form Parse", map[string]string{"error": "Data can't Handle"}))
		return
	}

	err = user.IsUserValid()
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	err = admin.userservice.Login(&user)
	if err != nil {
		web.RespondError(&w, err)
		return
	}

	if !user.IsEqualRole(adminRole) {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": "Invalid User for Role"}))
		return
	}

	token, err := admin.authcontroller.GetToken(&user)
	if err != nil {
		web.RespondError(&w, web.NewHTTPError("An error occur", http.StatusInternalServerError))
		return
	}
	web.RespondJSON(&w, http.StatusOK, models.TokenResponse{User: user, Token: token})

}

func (admin *AdminController) getAllUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	uid, err := admin.userservice.CheckUser(param, adminRole)
	if err != nil {
		web.RespondError(&w, err)
		return
	}

	users := []models.User{}
	err = admin.userservice.GetAll(uid, &users)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	web.RespondJSON(&w, http.StatusOK, users)
}

func (admin *AdminController) getUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	_, err := admin.userservice.CheckUser(param, adminRole)
	if err != nil {
		web.RespondError(&w, err)
		return
	}

	uid, err := web.ParseID(param["id"])
	if err != nil {
		// web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		web.RespondError(&w, web.NewValidationError("User ID", map[string]string{"error": "Invalid User ID"}))
		return
	}
	users := []models.User{}
	err = admin.userservice.Get(uid, &users)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	web.RespondJSON(&w, http.StatusOK, users)
}

func (admin *AdminController) addNewUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	_, err := admin.userservice.CheckUser(param, adminRole)
	if err != nil {
		web.RespondError(&w, err)
		return
	}

	user := *models.NewUserWithNewID()
	err = web.UnmarshalJSON(r, &user)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("Form Parse", map[string]string{"error": "Data can't handler"}))
		return
	}

	err = user.IsUserValid()
	if err != nil {
		web.RespondError(&w, err)
		return
	}

	err = admin.userservice.Register(&user)

	if err != nil {
		web.RespondError(&w, err)
		return
	}

	web.RespondJSON(&w, http.StatusOK, user.ID)
}

func (admin *AdminController) updateUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	_, err := admin.userservice.CheckUser(param, adminRole)
	if err != nil {
		web.RespondError(&w, err)
		return
	}

	uid, err := web.ParseID(param["id"])
	if err != nil {
		// web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		web.RespondError(&w, web.NewValidationError("User ID", map[string]string{"error": "Invalid User ID"}))
		return
	}
	user := models.User{}
	err = web.UnmarshalJSON(r, &user)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("Form Parse", map[string]string{"error": "Data can't Handle"}))
		return
	}

	err = user.IsUserValid()
	if err != nil {
		web.RespondError(&w, err)
		return
	}

	user.ID = *uid
	err = admin.userservice.Update(&user)

	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}
}

func (admin *AdminController) deleteUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	_, err := admin.userservice.CheckUser(param, adminRole)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	uid, err := web.ParseID(param["id"])
	if err != nil {
		// web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		web.RespondError(&w, web.NewValidationError("User ID", map[string]string{"error": "Invalid User ID"}))
		return
	}

	err = admin.userservice.Delete(uid)

	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}
}

func (admin *AdminController) allBookmark(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	_, err := admin.userservice.CheckUser(param, adminRole)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	userid, err := web.ParseID(param["id"])
	if err != nil {
		// web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		web.RespondError(&w, web.NewValidationError("User ID", map[string]string{"error": "Invalid User ID"}))
		return
	}
	bookmarks := []models.Bookmark{*models.NewBookmarkWithUserID(*userid)}
	// err = admin.bookmarkservice.GetAllBookmark(*userid, &bookmarks)
	// if err != nil {
	// 	web.RespondError(&w, err)
	// 	return
	// }

	web.RespondJSON(&w, http.StatusOK, bookmarks)
}

func (admin *AdminController) allCategory(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	_, err := admin.userservice.CheckUser(param, adminRole)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	userid, err := web.ParseID(param["id"])
	if err != nil {
		// web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		web.RespondError(&w, web.NewValidationError("User ID", map[string]string{"error": "Invalid User ID"}))
		return
	}
	categorys := []models.Category{*models.NewCategoryWithUserID(*userid)}
	// err = admin.categoryservice.GetAllBookmarkCategory(*userid, &categorys)
	// if err != nil {
	// 	web.RespondError(&w, err)
	// 	return
	// }

	web.RespondJSON(&w, http.StatusOK, categorys)
}
