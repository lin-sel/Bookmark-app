package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// Controller Structure
type Controller struct {
	authsrv *Authsrv
}

// Response Return As Successfull Login Response.
type Response struct {
	User
	Token string `json:"token"`
}

var secretKey = []byte("Private_Key")

// NewAuthController Return Controller Object
func NewAuthController(srv *Authsrv) *Controller {
	return &Controller{
		authsrv: srv,
	}
}

// RouterRgstr Register All Endpoint.
func (authcntrol *Controller) RouterRgstr(r *mux.Router) {
	r.HandleFunc("/api/user/register", authcntrol.registerUser).Methods("POST")
	r.HandleFunc("/api/user/login", authcntrol.login).Methods("POST")
}

func (authcntrol *Controller) registerUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}
	user := User{}

	if v := r.PostFormValue("name"); len(v) > 0 {
		user.Name = v
	}
	if len(user.Getname()) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("name required")
		return
	}
	if v := r.PostFormValue("username"); len(v) > 0 {
		user.Username = v
	}
	if len(user.Getusername()) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("username required")
		return
	}
	if v := r.PostFormValue("password"); len(v) > 0 {
		user.Password = v
	}

	if len(user.Getpassword()) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("password required")
		return
	}

	// Assign ID to User.
	user.ID = GetUUID()

	err = authcntrol.authsrv.Register(&user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	json.NewEncoder(w).Encode(user.ID)

}

func (authcntrol *Controller) login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Request")
		return
	}
	user := User{}
	if v := r.PostFormValue("username"); len(v) > 0 {
		user.Username = v
	}
	if v := r.PostFormValue("password"); len(v) > 0 {
		user.Password = v
	}

	err = authcntrol.authsrv.Login(&user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	authcntrol.GetToken(&user, &w)
}

// GetToken Return Token
func (authcntrol *Controller) GetToken(user *User, w *http.ResponseWriter) {
	// Create a claims map
	claims := jwt.MapClaims{
		"username": user.Getusername(),
		"IssuedAt": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		(*w).WriteHeader(http.StatusBadGateway)
		(*w).Write([]byte(err.Error()))
	}
	json.NewEncoder(*w).Encode(Response{Token: tokenString, User: *user})
}
