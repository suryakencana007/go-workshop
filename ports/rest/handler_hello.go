package rest

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"

	pkgHttp "go-workshop/pkg/http"
)


type Hello struct {
}

func (h *Hello) Register(ctx context.Context, router chi.Router) {
	router.Get("/hello", h.Hello)
	router.Get("/status-ok", h.StatusOk)
	router.Get("/status-accepted", h.StatusAccepted)
	router.Post("/upload", h.Upload)
}

func (h *Hello) Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	_ = pkgHttp.RequestJSONBody(w, r, http.StatusOK, map[string]interface{}{
		"Message": "Hello",
	})
}

func (h *Hello) StatusOk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	_ = pkgHttp.RequestJSONBody(w, r, http.StatusOK, map[string]interface{}{
		"Status": strconv.Itoa(http.StatusOK), 
		"Handler": "handler_status_ok",
	})
}

func (h *Hello) StatusAccepted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	_ = pkgHttp.RequestJSONBody(w, r, http.StatusAccepted, map[string]interface{}{
		"Status": strconv.Itoa(http.StatusAccepted), 
		"Handler": "handler_status_accepted",
	})
}

func (h *Hello) Upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	_ = pkgHttp.RequestJSONBody(w, r, http.StatusAccepted, map[string]interface{}{
		"Status": strconv.Itoa(http.StatusAccepted), 
		"Handler": "handler_upload",
	})
}