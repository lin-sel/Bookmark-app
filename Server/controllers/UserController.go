package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/lin-sel/bookmark-app/models"
	"github.com/lin-sel/bookmark-app/services"
	"github.com/lin-sel/bookmark-app/web"
)

// UserController Structure
type UserController struct {
	authsrv *services.UserService
}

// Response Return As Successfull Login Response.
type Response struct {
	models.User
	Token string `json:"token"`
}

// var secretKey = []byte("Private_Key")

// NewUserController Return UserController Instance
func NewUserController(srv *services.UserService) *UserController {
	return &UserController{
		authsrv: srv,
	}
}

// RouterRgstr Register All Endpoint.
func (authcntrol *UserController) RouterRgstr(r *mux.Router) {
	r.HandleFunc("/api/user/register", authcntrol.registerUser).Methods("POST")
	r.HandleFunc("/api/user/login", authcntrol.login).Methods("POST")
}

func (authcntrol *UserController) registerUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// err := r.ParseForm()
	user := models.User{}
	err := web.UnmarshalJSON(r, &user)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("Form Parse", map[string]string{"error": "Data can't handler"}))
		return
	}

	// if v := r.PostFormValue("name"); len(v) > 0 {
	// 	user.Name = v
	// }
	if len(user.Getname()) <= 0 {
		web.RespondError(&w, web.NewValidationError("require", map[string]string{"error": "name Required"}))
		return
	}
	// if v := r.PostFormValue("username"); len(v) > 0 {
	// 	user.Username = v
	// }
	if len(user.Getusername()) <= 0 {
		web.RespondError(&w, web.NewValidationError("require", map[string]string{"error": "username Required"}))
		return
	}
	// if v := r.PostFormValue("password"); len(v) > 0 {
	// 	user.Password = v
	// }

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

	json.NewEncoder(w).Encode(user.ID)

}

func (authcntrol *UserController) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// err := r.ParseForm()
	user := models.User{}
	err := web.UnmarshalJSON(r, &user)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("Form Parse", map[string]string{"error": "Data can't Handle"}))
		return
	}
	// if v := r.PostFormValue("username"); len(v) > 0 {
	// 	user.Username = v
	// }
	// if v := r.PostFormValue("password"); len(v) > 0 {
	// 	user.Password = v
	// }

	if len(user.Getusername()) == 0 {
		web.RespondError(&w, web.NewValidationError("Require", map[string]string{"error": "username required"}))
	}
	if len(user.Getpassword()) == 0 {
		web.RespondError(&w, web.NewValidationError("Require", map[string]string{"error": "password required"}))
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
		"IssuedAt": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		web.RespondError(w, web.NewValidationError("error", map[string]string{"msg": err.Error()}))
		return
	}
	json.NewEncoder(*w).Encode(Response{Token: tokenString, User: *user})
}
