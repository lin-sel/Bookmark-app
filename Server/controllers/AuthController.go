package controllers

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/mux"
	"github.com/lin-sel/bookmark-app/web"
	uuid "github.com/satori/go.uuid"
)

var secretKey = []byte("Private_Key")

// AuthUser Check For Valid User
func (cntrolr *Controller) AuthUser(h http.Handler) http.Handler {
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
			web.RespondError(&w, web.NewHTTPError("Access Denied.", http.StatusForbidden))
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
			h.ServeHTTP(w, r)
		} else {
			web.RespondError(&w, web.NewHTTPError(err.Error(), http.StatusForbidden))
		}
	})
}
