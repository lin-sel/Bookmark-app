package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/lin-sel/bookmark-app/models"
	"github.com/lin-sel/bookmark-app/services"
	"github.com/lin-sel/bookmark-app/web"
)

// Controller Structure
type Controller struct {
	bmsrv  *services.BookmarkService
	bmcsrv *services.BookmarkCategoryService
}

// NewController Return Bookmark Controller Object
func NewController(bookmarkservice *services.BookmarkService, bookmarkcategoryservice *services.BookmarkCategoryService) *Controller {
	return &Controller{
		bmsrv:  bookmarkservice,
		bmcsrv: bookmarkcategoryservice,
	}
}

// RouterRegstr Register All Endpoint of Bookmark to Router
func (cntrolr *Controller) RouterRegstr(r *mux.Router) {
	s := r.PathPrefix("/api/user/{userid}").Subrouter()
	s.Use(cntrolr.AuthUser)
	cntrolr.CategoryRgstr(s)
	s.HandleFunc("/bookmark", cntrolr.GetAllBookmark).Methods("GET")
	s.HandleFunc("/bookmark/{bookmarkid}", cntrolr.GetBookmarkByID).Methods("GET")
	s.HandleFunc("/bookmark/category/{bookmarkid}", cntrolr.GetBookmarkByCategory).Methods("GET")
	s.HandleFunc("/bookmark", cntrolr.AddBookmark).Methods("POST")
	s.HandleFunc("/bookmark/{bookmarkid}", cntrolr.UpdateBookmark).Methods("PUT")
	s.HandleFunc("/bookmark/{bookmarkid}", cntrolr.DeleteBookmark).Methods("DELETE")
}

// GetAllBookmark Return All Bookmark By UserID
func (cntrolr *Controller) GetAllBookmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["userid"]
	uid, err := web.ParseID(id)
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	bookmarks := []models.Bookmark{}
	err = cntrolr.bmsrv.GetAllBookmark(*uid, &bookmarks)
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusInternalServerError))
		return
	}
	json.NewEncoder(w).Encode(bookmarks)
}

// GetBookmarkByID Return Bookmark of Given ID
func (cntrolr *Controller) GetBookmarkByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	uid, err := web.ParseID(param["userid"])
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	var bid *uuid.UUID
	bid, err = web.ParseID(param["bookmarkid"])
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	bookmark := []models.Bookmark{}
	err = cntrolr.bmsrv.GetBookmark(*uid, *bid, &bookmark)
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusInternalServerError))
		return
	}
	json.NewEncoder(w).Encode(bookmark)
}

//GetBookmarkByCategory Return Bookmark By Category
func (cntrolr *Controller) GetBookmarkByCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	uid, err := web.ParseID(param["userid"])
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	var bid *uuid.UUID
	bid, err = web.ParseID(param["bookmarkid"])
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	bookmarks := []models.Bookmark{}
	err = cntrolr.bmsrv.GetBookmarkByCategory(*uid, *bid, &bookmarks)
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusInternalServerError))
		return
	}
	json.NewEncoder(w).Encode(bookmarks)
}

// UpdateBookmark Update Bookmark
func (cntrolr *Controller) UpdateBookmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	bookmark := models.Bookmark{}
	uid, err := web.ParseID(param["userid"])
	err = parseForm(&bookmark, r)
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	bookmark.UserID = *uid
	var id *uuid.UUID
	id, err = web.ParseID(param["bookmarkid"])
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError("Invalid ID", http.StatusBadRequest))
		return
	}
	bookmark.ID = *id

	err = cntrolr.bmsrv.UpdateBookmark(&bookmark)
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusInternalServerError))
		return
	}
}

// DeleteBookmark By ID
func (cntrolr *Controller) DeleteBookmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	uid, err := web.ParseID(param["userid"])
	err = r.ParseForm()
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	var id *uuid.UUID
	id, err = web.ParseID(param["bookmarkid"])
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError("Invalid ID", http.StatusBadRequest))
		return
	}

	err = cntrolr.bmsrv.DeleteBookmark(*uid, *id)
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusInternalServerError))
		return
	}
}

// AddBookmark Add New Data to Database
func (cntrolr *Controller) AddBookmark(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	uid, err := web.ParseID(mux.Vars(r)["userid"])
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError("Invalid User ID", http.StatusBadRequest))
		return
	}
	bookmark := models.Bookmark{}
	err = parseForm(&bookmark, r)
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	bookmark.UserID = *uid
	bookmark.ID = web.GetUUID()
	err = cntrolr.bmsrv.AddBookmark(&bookmark)
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusInternalServerError))
		return
	}
	json.NewEncoder(w).Encode(bookmark.ID)
}

func parseForm(bookmark *models.Bookmark, r *http.Request) error {
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
		id, err := web.ParseID(v)
		if err != nil {
			return errors.New("Invalid Category ID")
		}
		bookmark.CategoryID = *id
	}
	if (bookmark.GetCategoryID() == uuid.UUID{}) {
		return errors.New("categoryid is Required")
	}
	return nil
}
