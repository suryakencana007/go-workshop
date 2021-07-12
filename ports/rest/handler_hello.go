package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	pkgHttp "go-workshop/pkg/http"
)

type Hello struct {
}

func (h *Hello) Register(ctx context.Context, router chi.Router) {
	router.Get("/hello", h.Hello)
	router.Post("/upload", h.Upload)
	router.Get("/status-ok", h.Health)
	router.Get("/status-accepted", h.Accepted)
}

func (h *Hello) Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	_ = pkgHttp.RequestJSONBody(w, r, http.StatusOK, map[string]interface{}{
		"Message": "Hello",
	})
}

func (h *Hello) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pkgHttp.RequestJSONBody(w, r, http.StatusOK, map[string]interface{}{
		"Status":  "200",
		"Handler": "STATUS_OK",
	})
}

func (h *Hello) Accepted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pkgHttp.RequestJSONBody(w, r, http.StatusAccepted, map[string]interface{}{
		"Status":  "202",
		"Handler": "REQUEST_ACCEPTED",
	})
}

func (h *Hello) Upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	pkgHttp.RequestJSONBody(w, r, http.StatusAccepted, map[string]interface{}{
		"Status":  "202",
		"Handler": "REQUEST_SUCCESS",
	})
}
