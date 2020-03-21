package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/lin-sel/bookmark-app/models"
	"github.com/lin-sel/bookmark-app/services"
	"github.com/lin-sel/bookmark-app/web"
)

const session int64 = 600

// UserController Structure
type UserController struct {
	authsrv *services.UserService
}

// Response Return As Successfull Login Response.
type Response struct {
	models.User
	Token string `json:"token"`
}

// NewUserController Return UserController Instance
func NewUserController(srv *services.UserService) *UserController {
	return &UserController{
		authsrv: srv,
	}
}

// RouterRgstr Register All Endpoint.
func (authcntrol *UserController) RouterRgstr(r *mux.Router) {
	r.HandleFunc("/api/v1/user/register", authcntrol.registerUser).Methods("POST")
	r.HandleFunc("/api/v1/user/login", authcntrol.login).Methods("POST")
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
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"msg": err.Error()}))
		return
	}
	authcntrol.GetToken(&user, &w)
}

// GetToken Return Token
func (authcntrol *UserController) GetToken(user *models.User, w *http.ResponseWriter) {
	// Create a claims map
	fmt.Println(user.GetuserID())
	claims := jwt.MapClaims{
		"username": user.Getusername(),
		"userID":   user.GetuserID(),
		"IssuedAt": time.Now().Unix() + session,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		web.RespondError(w, web.NewValidationError("error", map[string]string{"msg": err.Error()}))
		return
	}
	web.RespondJSON(w, http.StatusOK, Response{Token: tokenString, User: *user})
}
