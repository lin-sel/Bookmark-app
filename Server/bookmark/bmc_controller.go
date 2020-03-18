package bookmark

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/lin-sel/bookmark-app/web"
)

// BMCController struct
type BMCController struct {
	srv *BMCService
}

// CategoryRgstr Register All Endpoint to Router
func (cntrlr *Controller) CategoryRgstr(s *mux.Router) {
	// s := r.PathPrefix("/api/user/{userid}/category").Subrouter()

	s.HandleFunc("/category", cntrlr.GetAllCategory).Methods("GET")
	s.HandleFunc("/category/{id}", cntrlr.GetCategoryByID).Methods("GET")
	s.HandleFunc("/category", cntrlr.AddCategory).Methods("POST")
	s.HandleFunc("/category/{id}", cntrlr.UpdateCategory).Methods("PUT")
	s.HandleFunc("/category/{id}", cntrlr.DeleteCategory).Methods("DELETE")
}

// GetAllCategory return All Category Of User
func (cntrlr *Controller) GetAllCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["userid"]
	uid, err := parseID(id)
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(w).Encode(err.Error())
		web.HeaderWrite(&w, http.StatusBadRequest, err)
		return
	}
	categories := []Category{}
	err = cntrlr.bmcsrv.GetAllBMCategory(*uid, &categories)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(err.Error())
		web.HeaderWrite(&w, http.StatusInternalServerError, err)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

// GetCategoryByID return All Category Of User
func (cntrlr *Controller) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	uid, err := parseID(param["userid"])
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(w).Encode(err.Error())
		web.HeaderWrite(&w, http.StatusBadRequest, err)
		return
	}
	var cid *uuid.UUID
	cid, err = parseID(param["id"])
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(w).Encode(err.Error())
		web.HeaderWrite(&w, http.StatusBadRequest, err)
		return
	}
	categories := []Category{}
	err = cntrlr.bmcsrv.GetBMCategory(*uid, *cid, &categories)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(err.Error())
		web.HeaderWrite(&w, http.StatusInternalServerError, err)
		return
	}
	json.NewEncoder(w).Encode(categories)

}

// AddCategory return All Category Of User
func (cntrlr *Controller) AddCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["userid"]
	uid, err := parseID(id)
	err = r.ParseForm()
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(w).Encode(err.Error())
		web.HeaderWrite(&w, http.StatusBadRequest, err)
		return
	}
	category := Category{}
	if v := r.PostFormValue("category"); len(v) > 0 {
		category.CName = v
	}
	if category.GetCategoryName() == "" {
		// w.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(w).Encode("category Require")
		web.HeaderWrite(&w, http.StatusBadRequest, errors.New("category Require"))
		return
	}
	category.UserID = *uid
	category.ID = GetUUID()

	err = cntrlr.bmcsrv.AddBMCategory(&category)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(err.Error())
		web.HeaderWrite(&w, http.StatusInternalServerError, err)
		return
	}

	json.NewEncoder(w).Encode(category.GetCategoryID)
}

// UpdateCategory Update Category
func (cntrlr *Controller) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	category := Category{}
	param := mux.Vars(r)

	uid, err := parseID(param["userid"])
	err = r.ParseForm()
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(w).Encode(err.Error())
		web.HeaderWrite(&w, http.StatusBadRequest, err)
		return
	}

	var id *uuid.UUID
	id, err = parseID(param["id"])
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(w).Encode("Invalid Id")
		web.HeaderWrite(&w, http.StatusBadRequest, errors.New("Invalid user ID"))
		return
	}
	category.ID = *id
	if v := r.PostFormValue("category"); len(v) > 0 {
		category.CName = v
	}
	if category.GetCategoryName() == "" {
		// w.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(w).Encode("category Require")
		web.HeaderWrite(&w, http.StatusBadRequest, errors.New("category Require"))
		return
	}
	category.UserID = *uid
	// category.ID = GetUUID()

	err = cntrlr.bmcsrv.UpdateBMCategory(&category)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(err.Error())
		web.HeaderWrite(&w, http.StatusInternalServerError, err)
		return
	}
}

// DeleteCategory Delete Category By ID
func (cntrlr *Controller) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	uid, err := parseID(param["userid"])
	err = r.ParseForm()
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(w).Encode(err.Error())
		web.HeaderWrite(&w, http.StatusBadRequest, err)
		return
	}

	var id *uuid.UUID
	id, err = parseID(param["id"])
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// json.NewEncoder(w).Encode("Invalid Id")
		web.HeaderWrite(&w, http.StatusBadRequest, errors.New("Invalid ID"))
		return
	}

	err = cntrlr.bmcsrv.DeleteBMCategory(*uid, *id)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// json.NewEncoder(w).Encode(err.Error())
		web.HeaderWrite(&w, http.StatusInternalServerError, err)
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

func parseID(id string) (*uuid.UUID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("Invalid User ID")
	}
	return &uid, nil
}
