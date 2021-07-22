package rest

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
	if r.Method == http.MethodPost {
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		f, fh, err := r.FormFile("upload-file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		dir, err := os.Getwd()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fileLocation := filepath.Join(dir, "files", fh.Filename)
		targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, fmt.Errorf("%s method not allowed", r.Method).Error(), http.StatusMethodNotAllowed)
		return
	}

	_ = pkgHttp.RequestJSONBody(w, r, http.StatusAccepted, payload)
}
