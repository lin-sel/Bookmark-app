package controllers

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lin-sel/bookmark-app/models"
	"github.com/lin-sel/bookmark-app/services"
	"github.com/lin-sel/bookmark-app/web"
	uuid "github.com/satori/go.uuid"
)

// Controller Structure
type Controller struct {
	bmsrv  *services.BookmarkService
	auth   *AuthController
	bmcsrv *services.BookmarkCategoryService
}

// NewController Return Bookmark Controller Object
func NewController(bookmarkservice *services.BookmarkService, bookmarkcategoryservice *services.BookmarkCategoryService, auth *AuthController) *Controller {
	return &Controller{
		bmsrv:  bookmarkservice,
		auth:   auth,
		bmcsrv: bookmarkcategoryservice,
	}
}

// RouterRegstr Register All Endpoint of Bookmark to Router
func (cntrolr *Controller) RouterRegstr(r *mux.Router) {
	s := r.PathPrefix("/{userid}").Subrouter()
	s.Use(cntrolr.auth.AuthUser)
	cntrolr.CategoryRgstr(s)
	s.HandleFunc("/bookmark", cntrolr.GetAllBookmark).Methods("GET")
	s.HandleFunc("/bookmark/{bookmarkid}", cntrolr.GetBookmarkByID).Methods("GET")
	s.HandleFunc("/bookmark/category/{categoryid}", cntrolr.GetBookmarkByCategory).Methods("GET")
	s.HandleFunc("/bookmark", cntrolr.AddBookmark).Methods("POST")
	s.HandleFunc("/bookmark/{bookmarkid}", cntrolr.UpdateBookmark).Methods("PUT")
	s.HandleFunc("/bookmark/{bookmarkid}", cntrolr.DeleteBookmark).Methods("DELETE")
}

// GetAllBookmark Return All Bookmark By UserID
func (cntrolr *Controller) GetAllBookmark(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["userid"]
	uid, err := web.ParseID(id)
	if err != nil {
		// web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		web.RespondError(&w, web.NewValidationError("User ID", map[string]string{"error": "Invalid User ID"}))
		return
	}
	bookmarks := []models.Bookmark{}
	err = cntrolr.bmsrv.GetAllBookmark(*uid, &bookmarks)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	web.RespondJSON(&w, http.StatusOK, bookmarks)
}

// GetBookmarkByID Return Bookmark of Given ID
func (cntrolr *Controller) GetBookmarkByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	uid, err := web.ParseID(param["userid"])
	if err != nil {
		// web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		web.RespondError(&w, web.NewValidationError("User ID", map[string]string{"error": "Invalid User ID"}))
		return
	}
	var bid *uuid.UUID
	bid, err = web.ParseID(param["bookmarkid"])
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"msg": err.Error()}))
		return
	}
	bookmark := []models.Bookmark{}
	err = cntrolr.bmsrv.GetBookmark(*uid, *bid, &bookmark)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	web.RespondJSON(&w, http.StatusOK, bookmark)
}

//GetBookmarkByCategory Return Bookmark By Category
func (cntrolr *Controller) GetBookmarkByCategory(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	uid, err := web.ParseID(param["userid"])
	if err != nil {
		web.RespondError(&w, web.NewValidationError("User ID", map[string]string{"error": "Invalid User ID"}))
		return
	}
	var bid *uuid.UUID
	bid, err = web.ParseID(param["categoryid"])
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"msg": "Invalid Category ID"}))
		return
	}
	bookmarks := []models.Bookmark{}
	err = cntrolr.bmsrv.GetBookmarkByCategory(*uid, *bid, &bookmarks)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	web.RespondJSON(&w, http.StatusOK, bookmarks)
}

// UpdateBookmark Update Bookmark
func (cntrolr *Controller) UpdateBookmark(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	bookmark := models.Bookmark{}
	uid, err := web.ParseID(param["userid"])
	err = web.UnmarshalJSON(r, &bookmark)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"msg": err.Error()}))
		return
	}
	err = validateBookmark(&bookmark)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"msg": err.Error()}))
		return
	}
	bookmark.UserID = *uid
	var id *uuid.UUID
	id, err = web.ParseID(param["bookmarkid"])
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"msg": err.Error()}))
		return
	}
	bookmark.ID = *id

	err = cntrolr.bmsrv.UpdateBookmark(&bookmark)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
}

// DeleteBookmark By ID
func (cntrolr *Controller) DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	uid, err := web.ParseID(param["userid"])
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"msg": err.Error()}))
		return
	}

	var id *uuid.UUID
	id, err = web.ParseID(param["bookmarkid"])
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"msg": err.Error()}))
		return
	}

	err = cntrolr.bmsrv.DeleteBookmark(*uid, *id)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
}

// AddBookmark Add New Data to Database
func (cntrolr *Controller) AddBookmark(w http.ResponseWriter, r *http.Request) {
	uid, err := web.ParseID(mux.Vars(r)["userid"])
	if err != nil {
		web.RespondError(&w, web.NewValidationError("User ID", map[string]string{"error": "Invalid User ID"}))
		return
	}
	bookmark := models.NewBookmarkWithID()
	err = web.UnmarshalJSON(r, bookmark)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"msg": err.Error()}))
		return
	}
	err = validateBookmark(bookmark)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"msg": err.Error()}))
		return
	}
	bookmark.UserID = *uid
	err = cntrolr.bmsrv.AddBookmark(bookmark)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	web.RespondJSON(&w, http.StatusOK, bookmark.ID)
}

func validateBookmark(bookmark *models.Bookmark) error {
	if bookmark.GetLabel() == "" {
		return errors.New("label is Required")
	}
	if bookmark.GetTag() == "" {
		return errors.New("tag is Required")
	}
	if bookmark.GetURL() == "" {
		return errors.New("url is Required")
	}
	if (bookmark.GetCategoryID() == uuid.UUID{}) {
		return errors.New("categoryid is Required")
	}
	return nil
}
