package controllers

import (
	"errors"
	"math"
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

// BookmarkRouteRegister Register All Endpoint of Bookmark to Router
func (cntrolr *Controller) BookmarkRouteRegister(r *mux.Router) {
	s := r.PathPrefix("/user/{userid}").Subrouter()
	s.Use(cntrolr.auth.AuthUser)
	cntrolr.CategoryRouteRegister(s)
	s.HandleFunc("/bookmark/{pagesize}/{pagenumber}", cntrolr.GetAllBookmark).Methods("GET")
	s.HandleFunc("/bookmark/{bookmarkid}", cntrolr.GetBookmarkByID).Methods("GET")
	s.HandleFunc("/bookmark/category/{categoryid}/{pagesize}/{pagenumber}", cntrolr.GetBookmarkByCategory).Methods("GET")
	s.HandleFunc("/bookmark", cntrolr.AddBookmark).Methods("POST")
	s.HandleFunc("/bookmark/{bookmarkid}", cntrolr.UpdateBookmark).Methods("PUT")
	s.HandleFunc("/bookmark/{bookmarkid}", cntrolr.DeleteBookmark).Methods("DELETE")
}

// GetAllBookmark Return All Bookmark By UserID
func (cntrolr *Controller) GetAllBookmark(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	uid, err := web.ParseID(param["userid"])
	if err != nil {
		// web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		web.RespondError(&w, web.NewValidationError("User ID", map[string]string{"error": "Invalid User ID"}))
		return
	}
	pagesize := web.ParseInt64(param["pagesize"])
	if *pagesize == int64(0) {
		*pagesize = 100
	}
	pagenumber := web.ParseInt64(param["pagenumber"])
	if *pagesize == int64(0) {
		*pagesize = 1
	}
	response := models.NewResponseBookmark(&[]models.Bookmark{*models.NewBookmarkWithUserID(*uid)}, *pagenumber, *pagesize)
	err = cntrolr.bmsrv.GetAllBookmark(*uid, response)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	var count int64
	err = cntrolr.bmsrv.GetTotalCount(models.NewBookmarkWithUserID(*uid), &count, "user_id", *uid)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	response.TotalPage = int64(math.Ceil((float64(count) / float64(*pagesize))))
	response.TotalCount = count
	if response.TotalPage == 0 && count > 0 {
		response.TotalPage = 1
	}
	web.RespondJSON(&w, http.StatusOK, response)
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
	bookmark := []models.Bookmark{models.Bookmark{}}
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
	var cid *uuid.UUID
	cid, err = web.ParseID(param["categoryid"])
	if err != nil {
		web.RespondError(&w, web.NewValidationError("error", map[string]string{"msg": "Invalid Category ID"}))
		return
	}
	pagesize := web.ParseInt64(param["pagesize"])
	if *pagesize == int64(0) {
		*pagesize = 100
	}
	pagenumber := web.ParseInt64(param["pagenumber"])
	if *pagesize == int64(0) {
		*pagesize = 1
	}
	response := models.NewResponseBookmark(&[]models.Bookmark{*models.NewBookmarkWithUserID(*uid)}, *pagenumber, *pagesize)
	err = cntrolr.bmsrv.GetBookmarkByCategory(*cid, response)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	var count int64
	err = cntrolr.bmsrv.GetTotalCount(models.NewBookmarkWithUserID(*uid), &count, "category_id", *cid)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	response.TotalPage = int64(math.Ceil((float64(count) / float64(*pagesize))))
	response.TotalCount = count
	if response.TotalPage == 0 && count > 0 {
		response.TotalPage = 1
	}
	web.RespondJSON(&w, http.StatusOK, response)
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
	bookmark := models.NewBookmarkWithUserID(*uid)
	err = cntrolr.bmsrv.DeleteBookmark(*id, bookmark)
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
