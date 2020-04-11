package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lin-sel/bookmark-app/models"
	"github.com/lin-sel/bookmark-app/services"
	"github.com/lin-sel/bookmark-app/web"
)

const session int64 = 3600

const userRole = "user"

// response Return After Successful Login
type useresponse struct {
	models.User
	Token string `json:"token"`
}

// UserController Structure
type UserController struct {
	auth    *AuthController
	authsrv *services.UserService
}

// NewUserController Return UserController Instance
func NewUserController(srv *services.UserService, auth *AuthController) *UserController {
	return &UserController{
		auth:    auth,
		authsrv: srv,
	}
}

// UserRouteRegister Register All Endpoint.
func (authcntrol *UserController) UserRouteRegister(r *mux.Router) {
	r.HandleFunc("/user/register", authcntrol.registerUser).Methods("POST")
	r.HandleFunc("/user/login", authcntrol.login).Methods("POST")
	s := r.PathPrefix("/user/{userid}/user").Subrouter()
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

	err = user.IsUserValid()
	if err != nil {
		web.RespondError(&w, err)
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

	// if len(user.Getusername()) == 0 {
	// 	web.RespondError(&w, web.NewValidationError("Require", map[string]string{"error": "username required"}))
	// 	return
	// }
	// if len(user.Getpassword()) == 0 {
	// 	web.RespondError(&w, web.NewValidationError("Require", map[string]string{"error": "password required"}))
	// 	return
	// }

	err = user.IsUserValid()
	if err != nil {
		web.RespondError(&w, err)
		return
	}

	err = authcntrol.authsrv.Login(&user)

	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}
	if !user.IsEqualRole(userRole) {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": "Invalid User for Role"}))
		return
	}
	token, err := authcntrol.auth.GetToken(&user)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}

	web.RespondJSON(&w, http.StatusOK, models.TokenResponse{Token: token, User: user})
}

func (authcntrol *UserController) update(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	uid, err := authcntrol.authsrv.CheckUser(param, userRole)
	if err != nil {
		web.RespondError(&w, err)
		return
	}

	user := models.User{Role: "user"}
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
	err = authcntrol.authsrv.Update(&user)

	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}
}

func (authcntrol *UserController) delete(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	uid, err := authcntrol.authsrv.CheckUser(param, userRole)
	if err != nil {
		web.RespondError(&w, err)
		return
	}

	err = authcntrol.authsrv.Delete(uid)

	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}
}

func (authcntrol *UserController) get(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	uid, err := authcntrol.authsrv.CheckUser(param, userRole)
	if err != nil {
		web.RespondError(&w, err)
		return
	}

	user := []models.User{*models.NewUserWithID(*uid)}
	err = authcntrol.authsrv.Get(uid, &user)

	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"error": err.Error()}))
		return
	}

	web.RespondJSON(&w, http.StatusOK, user)
}
