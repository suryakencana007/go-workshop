package rest

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

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

	fmt.Println("starting upload...")

	// HTTP method check
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Errorf("%s method not allowed", r.Method).Error(), http.StatusMethodNotAllowed)
		return
	}

	// 32MB max upload size
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// get value
	v := r.PostFormValue("file-name")

	// get file
	f, fh, err := r.FormFile("upload-file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// close the file when finished
	defer f.Close()

	// file info log
	fmt.Printf("File Name\t: %+v\n", v)
	fmt.Printf("Uploaded File\t: %+v\n", fh.Filename)
	fmt.Printf("File Size\t: %+v\n", fh.Size)
	fmt.Printf("MIME Header\t: %+v\n", fh.Header)

	// get current path
	dir, err := os.Getwd()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fileLocation := filepath.Join(dir, "files")
	fl := filepath.FromSlash(fileLocation)
	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile(fl, fmt.Sprintf("*.%s", filepath.Ext(fh.Filename)))
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

	fmt.Println("upload finished")

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"Status":  "202",
		"Handler": "UPLOAD_SUCCESS",
	}
	pkgHttp.RequestJSONBody(w, r, http.StatusAccepted, response)

}
