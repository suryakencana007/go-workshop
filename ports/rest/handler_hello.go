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
	payload := map[string]interface{}{
		"Status":  strconv.Itoa(http.StatusOK),
		"Handler": "Health",
	}
	_ = pkgHttp.RequestJSONBody(w, r, http.StatusOK, payload)
}

func (h *Hello) Accepted(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	payload := map[string]interface{}{
		"Status":  strconv.Itoa(http.StatusAccepted),
		"Handler": "Accepted",
	}
	_ = pkgHttp.RequestJSONBody(w, r, http.StatusAccepted, payload)
}

func (h *Hello) Upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	payload := map[string]interface{}{
		"Status":  strconv.Itoa(http.StatusAccepted),
		"Handler": "Upload",
	}
	_ = pkgHttp.RequestJSONBody(w, r, http.StatusAccepted, payload)
}
