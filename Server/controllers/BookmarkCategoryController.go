package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lin-sel/bookmark-app/models"
	"github.com/lin-sel/bookmark-app/web"
	uuid "github.com/satori/go.uuid"
)

// CategoryRgstr Register All Endpoint to Router
func (cntrlr *Controller) CategoryRgstr(s *mux.Router) {
	s.HandleFunc("/category", cntrlr.GetAllCategory).Methods("GET")
	s.HandleFunc("/category/{categoryid}", cntrlr.GetCategoryByID).Methods("GET")
	s.HandleFunc("/category", cntrlr.AddCategory).Methods("POST")
	s.HandleFunc("/category/{categoryid}", cntrlr.UpdateCategory).Methods("PUT")
	s.HandleFunc("/category/{categoryid}", cntrlr.DeleteCategory).Methods("DELETE")
}

// GetAllCategory return All Category Of User
func (cntrlr *Controller) GetAllCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["userid"]
	uid, err := web.ParseID(id)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("Invalid User ID", map[string]string{"error": "Invalid User ID"}))
		return
	}
	categories := []models.Category{}
	err = cntrlr.bmcsrv.GetAllBookmarkCategory(*uid, &categories)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	fmt.Println(categories[0].GetCategoryName(), "Called")
	web.RespondJSON(&w, http.StatusOK, categories)
}

// GetCategoryByID return All Category Of User
func (cntrlr *Controller) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	uid, err := web.ParseID(param["userid"])
	if err != nil {
		web.RespondError(&w, web.NewValidationError("User ID", map[string]string{"error": "Invalid User ID"}))
		return
	}
	var cid *uuid.UUID
	cid, err = web.ParseID(param["categoryid"])
	if err != nil {
		web.RespondError(&w, web.NewValidationError("require", map[string]string{"error": "Category ID Required"}))
		return
	}
	categories := []models.Category{}
	err = cntrlr.bmcsrv.GetBookmarkCategory(*uid, *cid, &categories)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	web.RespondJSON(&w, http.StatusOK, categories)
}

// AddCategory return All Category Of User
func (cntrlr *Controller) AddCategory(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["userid"]
	uid, err := web.ParseID(id)
	category := models.Category{}
	err = web.UnmarshalJSON(r, &category)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("Form Parse", map[string]string{"error": "data can't handle"}))
		return
	}

	if category.GetCategoryName() == "" {
		web.RespondError(&w, web.NewValidationError("Require", map[string]string{"error": "Category Name Required"}))
		return
	}
	category.UserID = *uid
	category.ID = web.GetUUID()
	err = cntrlr.bmcsrv.AddBookmarkCategory(&category)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
	fmt.Println(category.GetCategoryID())
	web.RespondJSON(&w, http.StatusOK, category.GetCategoryID())
}

// UpdateCategory Update Category
func (cntrlr *Controller) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	category := models.Category{}
	param := mux.Vars(r)

	uid, err := web.ParseID(param["userid"])
	if err != nil {
		web.RespondError(&w, web.NewValidationError("Form Parse", map[string]string{"error": "data can't handle"}))
		return
	}

	var id *uuid.UUID
	id, err = web.ParseID(param["categoryid"])
	if err != nil {
		web.RespondError(&w, web.NewValidationError("User ID", map[string]string{"error": "Invalid User ID"}))
		return
	}
	category.ID = *id
	err = web.UnmarshalJSON(r, &category)
	if err != nil {
		web.RespondError(&w, web.NewValidationError("Error", map[string]string{"msg": "Data can't handle"}))
		return
	}

	if category.GetCategoryName() == "" {
		web.RespondError(&w, web.NewValidationError("require", map[string]string{"error": "Category ID Required"}))
		return
	}
	category.UserID = *uid
	err = cntrlr.bmcsrv.UpdateBookmarkCategory(&category)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
}

// DeleteCategory Delete Category By ID
func (cntrlr *Controller) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	uid, err := web.ParseID(param["userid"])
	if err != nil {
		web.RespondError(&w, web.NewValidationError("User ID", map[string]string{"error": err.Error()}))
		return
	}

	var id *uuid.UUID
	id, err = web.ParseID(param["categoryid"])
	if err != nil {
		web.RespondError(&w, web.NewValidationError("Category ID", map[string]string{"error": "Invalid Category ID"}))
		return
	}

	err = cntrlr.bmcsrv.DeleteBookmarkCategory(*uid, *id)
	if err != nil {
		web.RespondError(&w, err)
		return
	}
}

// RecentCategory Return Recent Added or Modified Category.
func (cntrlr *Controller) RecentCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
