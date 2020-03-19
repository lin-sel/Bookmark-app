package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/lin-sel/bookmark-app/models"
	"github.com/lin-sel/bookmark-app/web"
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
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["userid"]
	uid, err := web.ParseID(id)
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError("Invalid User ID", http.StatusBadRequest))
		return
	}
	categories := []models.Category{}
	err = cntrlr.bmcsrv.GetAllBookmarkCategory(*uid, &categories)
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusInternalServerError))
		return
	}
	json.NewEncoder(w).Encode(categories)
}

// GetCategoryByID return All Category Of User
func (cntrlr *Controller) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	uid, err := web.ParseID(param["userid"])
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	var cid *uuid.UUID
	cid, err = web.ParseID(param["categoryid"])
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	categories := []models.Category{}
	err = cntrlr.bmcsrv.GetBookmarkCategory(*uid, *cid, &categories)
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusInternalServerError))
		return
	}
	json.NewEncoder(w).Encode(categories)

}

// AddCategory return All Category Of User
func (cntrlr *Controller) AddCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["userid"]
	uid, err := web.ParseID(id)
	err = r.ParseForm()
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}
	category := models.Category{}
	if v := r.PostFormValue("category"); len(v) > 0 {
		category.CName = v
	}
	if category.GetCategoryName() == "" {
		web.WriteErrorResponse(&w, web.NewHTTPError("category Require", http.StatusBadRequest))
		return
	}
	category.UserID = *uid
	category.ID = web.GetUUID()
	err = cntrlr.bmcsrv.AddBookmarkCategory(&category)
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusInternalServerError))
		return
	}

	json.NewEncoder(w).Encode(category.GetCategoryID())
}

// UpdateCategory Update Category
func (cntrlr *Controller) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	category := models.Category{}
	param := mux.Vars(r)

	uid, err := web.ParseID(param["userid"])
	err = r.ParseForm()
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	var id *uuid.UUID
	id, err = web.ParseID(param["categoryid"])
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError("Invalid user ID", http.StatusBadRequest))
		return
	}
	category.ID = *id
	if v := r.PostFormValue("category"); len(v) > 0 {
		category.CName = v
	}
	if category.GetCategoryName() == "" {
		web.WriteErrorResponse(&w, web.NewHTTPError("category required", http.StatusBadRequest))
		return
	}
	category.UserID = *uid
	// category.ID = GetUUID()

	err = cntrlr.bmcsrv.UpdateBookmarkCategory(&category)
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusInternalServerError))
		return
	}
}

// DeleteCategory Delete Category By ID
func (cntrlr *Controller) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	uid, err := web.ParseID(param["userid"])
	err = r.ParseForm()
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusBadRequest))
		return
	}

	var id *uuid.UUID
	id, err = web.ParseID(param["categoryid"])
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError("Invalid ID", http.StatusBadRequest))
		return
	}

	err = cntrlr.bmcsrv.DeleteBookmarkCategory(*uid, *id)
	if err != nil {
		web.WriteErrorResponse(&w, web.NewHTTPError(err.Error(), http.StatusInternalServerError))
		return
	}
}

// RecentCategory Return Recent Added or Modified Category.
func (cntrlr *Controller) RecentCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

// func (cntrlr *BMCController)ValidateUser(id string) *uuid.UUID, error{
// 	uid, err := uuid.Parse(id)
// 	if err != nil {
// 		return nil, errors.New("Invalid User")
// 	}
// }
