package controllers

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	"github.com/lin-sel/bookmark-app/models"
	"github.com/lin-sel/bookmark-app/web"
	uuid "github.com/satori/go.uuid"
)

var secretKey []byte

// AuthController Type
type AuthController struct{}

// NewAuthController return New Instance.
func NewAuthController(key []byte) *AuthController {
	secretKey = key
	return &AuthController{}
}

// AuthUser Check For Valid User
func (cntrolr *AuthController) AuthUser(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := request.HeaderExtractor{"token"}.ExtractToken(r)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return secretKey, nil
		})
		if err != nil {
			fmt.Println(err)
			web.RespondError(&w, web.NewHTTPError("Hello Access Denied.", http.StatusForbidden))
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var apiuserid, userid *uuid.UUID
			var err error
			apiuserid, err = web.ParseID(mux.Vars(r)["userid"])
			if err != nil {
				web.RespondError(&w, web.NewHTTPError(err.Error(), http.StatusForbidden))
				return
			}
			id, er := claims["userID"].(string)
			if !er {
				web.RespondError(&w, web.NewHTTPError("nil", http.StatusForbidden))
				return
			}
			userid, err = web.ParseID(id)
			if err != nil {
				web.RespondError(&w, web.NewHTTPError(err.Error(), http.StatusForbidden))
				return
			}

			if *userid != *apiuserid {
				web.RespondError(&w, web.NewHTTPError("Access Denied.", http.StatusForbidden))
				return
			}

			issueAt, er := claims["IssuedAt"].(float64)
			if issueAt < float64(time.Now().Unix()) {
				web.RespondError(&w, web.NewHTTPError("Session Expire Please Login Again.", http.StatusForbidden))
				return
			}
			if r.Method == "OPTION" {
				w.WriteHeader(http.StatusOK)
				return
			}
			h.ServeHTTP(w, r)
		} else {
			web.RespondError(&w, web.NewHTTPError(err.Error(), http.StatusForbidden))
		}
	})
}

// GetToken Return Token
func (cntrolr *AuthController) GetToken(user *models.User, w *http.ResponseWriter) (string, error) {
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
		return "", err
	}
	return tokenString, nil
}
