package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lin-sel/bookmark-app/models"
	"github.com/lin-sel/bookmark-app/services"
	"github.com/lin-sel/bookmark-app/web"
)

const session int64 = 600

// UserController Structure
type UserController struct {
	auth    *AuthController
	authsrv *services.UserService
}

// Response Return As Successfull Login Response.
type Response struct {
	models.User
	Token string `json:"token"`
}

// NewUserController Return UserController Instance
func NewUserController(srv *services.UserService, auth *AuthController) *UserController {
	return &UserController{
		auth:    auth,
		authsrv: srv,
	}
}

// RouterRgstr Register All Endpoint.
func (authcntrol *UserController) RouterRgstr(r *mux.Router) {
	r.HandleFunc("/register", authcntrol.registerUser).Methods("POST")
	r.HandleFunc("/login", authcntrol.login).Methods("POST")
	s := r.PathPrefix("/{userid}/user").Subrouter()
	s.Use(authcntrol.auth.AuthUser)
	s.HandleFunc("", authcntrol.get).Methods("GET")
	s.HandleFunc("", authcntrol.delete).Methods("DELETE")
	s.HandleFunc("", authcntrol.update).Methods("PUT")
}

func (authcntrol *UserController) registerUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
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

	// Assign ID to User.
	user.ID = web.GetUUID()

	err = authcntrol.authsrv.Register(&user)

	if err != nil {
		web.RespondError(&w, err)
		return
	}

	web.RespondJSON(&w, http.StatusOK, user.ID)

}

func (authcntrol *UserController) login(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
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

	err = authcntrol.authsrv.Login(&user)

	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}
	token, err := authcntrol.auth.GetToken(&user, &w)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}

	web.RespondJSON(&w, http.StatusOK, Response{Token: token, User: user})
}

func (authcntrol *UserController) update(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	uid, err := web.ParseID(param["userid"])
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

	err = authcntrol.authsrv.Update(&user)

	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}
}

func (authcntrol *UserController) delete(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	uid, err := web.ParseID(param["userid"])
	if err != nil {
		// web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		web.RespondError(&w, web.NewValidationError("User ID", map[string]string{"error": "Invalid User ID"}))
		return
	}

	err = authcntrol.authsrv.Delete(uid)

	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}
}

func (authcntrol *UserController) get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Called")
	param := mux.Vars(r)
	uid, err := web.ParseID(param["userid"])
	if err != nil {
		// web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		web.RespondError(&w, web.NewValidationError("User ID", map[string]string{"error": "Invalid User ID"}))
		return
	}

	user := models.User{}
	err = authcntrol.authsrv.Get(uid, &user)

	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}

	user.Category = []models.Category{}
	web.RespondJSON(&w, http.StatusOK, user)
}
