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
	authsrv *services.Authsrv
}

// Response Return As Successfull Login Response.
type Response struct {
	models.User
	Token string `json:"token"`
}

// var secretKey = []byte("Private_Key")

// NewUserController Return UserController Instance
func NewUserController(srv *services.Authsrv) *UserController {
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
	err := r.ParseForm()
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError("Invalid Request", http.StatusBadRequest))
		return
	}
	user := models.User{}

	if v := r.PostFormValue("name"); len(v) > 0 {
		user.Name = v
	}
	if len(user.Getname()) <= 0 {
		web.WriteErrorResponse(&w, web.NewHTTPError("name Required", http.StatusBadRequest))
		return
	}
	if v := r.PostFormValue("username"); len(v) > 0 {
		user.Username = v
	}
	if len(user.Getusername()) <= 0 {
		web.WriteErrorResponse(&w, web.NewHTTPError("username Required", http.StatusBadRequest))
		return
	}
	if v := r.PostFormValue("password"); len(v) > 0 {
		user.Password = v
	}

	if len(user.Getpassword()) <= 0 {
		web.WriteErrorResponse(&w, web.NewHTTPError("password required", http.StatusBadRequest))
		return
	}

	// Assign ID to User.
	user.ID = web.GetUUID()

	err = authcntrol.authsrv.Register(&user)

	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusInternalServerError))

		return
	}

	json.NewEncoder(w).Encode(user.ID)

}

func (authcntrol *UserController) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError("Invalid Request", http.StatusBadRequest))
		return
	}
	user := models.User{}
	if v := r.PostFormValue("username"); len(v) > 0 {
		user.Username = v
	}
	if v := r.PostFormValue("password"); len(v) > 0 {
		user.Password = v
	}

	err = authcntrol.authsrv.Login(&user)

	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))

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
		web.WriteErrorResponse(w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	json.NewEncoder(*w).Encode(Response{Token: tokenString, User: *user})
}
