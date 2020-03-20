package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lin-sel/bookmark-app/web"
)

// HealthController provides method to check health and readiness
type HealthController struct {
}

// NewHealthController returns a new instance of HealthController
func NewHealthController() *HealthController {
	return &HealthController{}
}

// RegisterRoutes implements interface RouteSpecifier
func (controller *Controller) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/api/health", controller.healthCheck).Methods("GET")
}

func (controller *Controller) healthCheck(w http.ResponseWriter, r *http.Request) {
	web.RespondJSON(&w, http.StatusOK, `{"msg":"OK"}`)
}
