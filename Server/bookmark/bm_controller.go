package bookmark

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Controller Structure
type Controller struct {
	bmsrv  *BMService
	bmcsrv *BMCService
}

var secretKey = []byte("Private_Key")

// NewController Return Bookmark Controller Object
func NewController(bmservice *BMService, bmcservice *BMCService) *Controller {
	return &Controller{
		bmsrv:  bmservice,
		bmcsrv: bmcservice,
	}
}

// RouterRegstr Register All Endpoint of Bookmark to Router
func (cntrolr *Controller) RouterRegstr(r *mux.Router) {
	s := r.PathPrefix("/api/user/{userid}").Subrouter()
	s.Use(cntrolr.AuthUser)
	cntrolr.CategoryRgstr(s)
	s.HandleFunc("/bookmark", cntrolr.GetAllBookmark).Methods("GET")
	s.HandleFunc("/bookmark/{id}", cntrolr.GetBookmarkByID).Methods("GET")
	s.HandleFunc("/bookmark/category/{id}", cntrolr.GetBookmarkByCategory).Methods("GET")
	s.HandleFunc("/bookmark", cntrolr.AddBookmark).Methods("POST")
	s.HandleFunc("/bookmark/{id}", cntrolr.UpdateBookmark).Methods("PUT")
	s.HandleFunc("/bookmark/{id}", cntrolr.DeleteBookmark).Methods("DELETE")
}

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
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Access Denied; Please Login First."))
			return
		}
		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// // If token is valid
			// response := make(map[string]string)
			// // response["user"] = claims["username"]
			// response["time"] = time.Now().String()
			// response["user"] = claims["username"].(string)
			// responseJSON, _ := json.Marshal(response)
			// w.Write(responseJSON)
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(err.Error()))
		}
	})
}

// GetAllBookmark Return All Bookmark By UserID
func (cntrolr *Controller) GetAllBookmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["userid"]
	uid, err := parseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	bookmarks := []Bookmark{}
	err = cntrolr.bmsrv.GetAllBookmark(*uid, &bookmarks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(bookmarks)
}

// GetBookmarkByID Return Bookmark of Given ID
func (cntrolr *Controller) GetBookmarkByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	uid, err := parseID(param["userid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	var bid *uuid.UUID
	bid, err = parseID(param["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	bookmark := []Bookmark{}
	err = cntrolr.bmsrv.GetBookmark(*uid, *bid, &bookmark)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(bookmark)
}

//GetBookmarkByCategory Return Bookmark By Category
func (cntrolr *Controller) GetBookmarkByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	uid, err := parseID(param["userid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	var bid *uuid.UUID
	bid, err = parseID(param["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	bookmarks := []Bookmark{}
	err = cntrolr.bmsrv.GetBookmarkByCategory(*uid, *bid, &bookmarks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(bookmarks)
}

// UpdateBookmark Update Bookmark
func (cntrolr *Controller) UpdateBookmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	bookmark := Bookmark{}
	uid, err := parseID(param["userid"])
	err = parseForm(&bookmark, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	bookmark.UserID = *uid
	var id *uuid.UUID
	id, err = parseID(param["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Id")
		return
	}
	bookmark.ID = *id

	err = cntrolr.bmsrv.UpdateBookmark(&bookmark)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
}

// DeleteBookmark By ID
func (cntrolr *Controller) DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	uid, err := parseID(param["userid"])
	err = r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var id *uuid.UUID
	id, err = parseID(param["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Id")
		return
	}

	err = cntrolr.bmsrv.DeleteBookmark(*uid, *id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
}

// AddBookmark Add New Data to Database
func (cntrolr *Controller) AddBookmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	uid, err := parseID(mux.Vars(r)["userid"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid User ID")
		return
	}
	bookmark := Bookmark{}
	err = parseForm(&bookmark, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	bookmark.UserID = *uid
	bookmark.ID = GetUUID()
	err = cntrolr.bmsrv.AddBookmark(&bookmark)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(bookmark.ID)
}

func parseForm(bookmark *Bookmark, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}
	if v := r.PostFormValue("label"); len(v) > 0 {
		bookmark.Label = v
	}
	if bookmark.GetLabel() == "" {
		return errors.New("label is Required")
	}
	if v := r.PostFormValue("tag"); len(v) > 0 {
		bookmark.Tag = v
	}
	if bookmark.GetTag() == "" {
		return errors.New("tag is Required")
	}
	if v := r.PostFormValue("url"); len(v) > 0 {
		bookmark.URL = v
	}
	if bookmark.GetURL() == "" {
		return errors.New("url is Required")
	}
	if v := r.PostFormValue("categoryid"); len(v) > 0 {
		id, err := parseID(v)
		if err != nil {
			return errors.New("Invalid Category ID")
		}
		bookmark.CategoryID = *id
	}
	// fmt.Println(bookmark.GetCategoryID() == uuid.UUID{})
	if (bookmark.GetCategoryID() == uuid.UUID{}) {
		return errors.New("categoryid is Required")
	}
	return nil
}
