package rest

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

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
	fmt.Println("upload")
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Errorf("%s method not allowed", r.Method).Error(), http.StatusMethodNotAllowed)
		return
	}
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

	fmt.Printf("Uploaded File: %+v\n", fh.Filename)
	fmt.Printf("File Size: %+v\n", fh.Size)
	fmt.Printf("MIME Header: %+v\n", fh.Header)

	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileLocation := filepath.Join(dir, "files")
	fl := filepath.FromSlash(fileLocation)
	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile(fl, fmt.Sprintf("*.%s",  filepath.Ext(fh.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// write this byte array to our temporary file
	tempFile.Write(fileBytes)

	r.ParseForm()
	for key, val := range r.Form {
		payload[key] = val[0]
	}
	_ = pkgHttp.RequestJSONBody(w, r, http.StatusAccepted, payload)
}
