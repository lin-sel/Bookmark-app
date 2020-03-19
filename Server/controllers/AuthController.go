package controllers

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/lin-sel/bookmark-app/web"
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
			web.WriteErrorResponse(&w, web.NewHTTPError("Access Denied.", http.StatusForbidden))
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// // If token is valid
			// response := make(map[string]string)
			// // response["user"] = claims["username"]
			// response["time"] = time.Now().String()
			// response["user"] = claims["username"].(string)
			// responseJSON, _ := json.Marshal(response)
			// w.Write(responseJSON)
			var apiuserid, userid *uuid.UUID
			var err error
			apiuserid, err = web.ParseID(mux.Vars(r)["userid"])
			if err != nil {
				web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusForbidden))
				return
			}
			id, er := claims["userID"].(string)
			if !er {
				web.WriteErrorResponse(&w, web.NewHTTPError("nil", http.StatusForbidden))
				return
			}
			userid, err = web.ParseID(id)
			if err != nil {
				web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusForbidden))
				return
			}

			if *userid != *apiuserid {
				web.WriteErrorResponse(&w, web.NewHTTPError("Access Denied.", http.StatusForbidden))
				return
			}
			h.ServeHTTP(w, r)
		} else {
			web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusForbidden))
		}
	})
}