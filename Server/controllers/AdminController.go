package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lin-sel/bookmark-app/models"
	"github.com/lin-sel/bookmark-app/services"
	"github.com/lin-sel/bookmark-app/web"
)

// response Return After Successful Login
type response struct {
	models.Admin
	Token string `json:"token"`
}

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
}

// login Admin Login
func (admin *AdminController) login(w http.ResponseWriter, r *http.Request) {
	user := models.Admin{}
	err := web.UnmarshalJSON(r, &user)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("Form Parse", map[string]string{"error": "Data can't Handle"}))
		return
	}

	if len(user.Getusername()) == 0 {
		web.RespondError(&w, web.NewValidationError("Require", map[string]string{"error": "username required"}))
		return
	}
	if len(user.Getpassword()) == 0 {
		web.RespondError(&w, web.NewValidationError("Require", map[string]string{"error": "password required"}))
		return
	}
	fmt.Println(user)
	err = admin.adminservice.Login(&user)
	if err != nil {
		web.RespondError(&w, err)
		return
	}

	token, err := admin.authcontroller.GetToken(&user)
	if err != nil {
		web.RespondError(&w, web.NewHTTPError("An error occur", http.StatusInternalServerError))
		return
	}
	web.RespondJSON(&w, http.StatusOK, response{Admin: user, Token: token})

}

func (admin *AdminController) getAllUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["userid"]
	uid, err := web.ParseID(id)
	if err != nil {
		// web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		web.RespondError(&w, web.NewValidationError("User ID", map[string]string{"error": "Invalid User ID"}))
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
	ids := mux.Vars(r)
	_, err := web.ParseID(ids["userid"])
	if err != nil {
		// web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		web.RespondError(&w, web.NewValidationError("User ID", map[string]string{"error": "Invalid Admin ID"}))
		return
	}
	uid, err := web.ParseID(ids["id"])
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
	user := *models.NewUserWithNewID()
	err := web.UnmarshalJSON(r, &user)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("Form Parse", map[string]string{"error": "Data can't handler"}))
		return
	}

	if len(user.Getname()) <= 0 {
		web.RespondError(&w, web.NewValidationError("require", map[string]string{"error": "name Required"}))
		return
	}

	if len(user.Getusername()) <= 0 {
		web.RespondError(&w, web.NewValidationError("require", map[string]string{"error": "username Required"}))
		return
	}

	if len(user.Getpassword()) <= 0 {
		web.RespondError(&w, web.NewValidationError("require", map[string]string{"error": "password Required"}))
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

	if len(user.Getusername()) == 0 {
		web.RespondError(&w, web.NewValidationError("Require", map[string]string{"error": "username required"}))
		return
	}
	if len(user.Getpassword()) == 0 {
		web.RespondError(&w, web.NewValidationError("Require", map[string]string{"error": "password required"}))
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
