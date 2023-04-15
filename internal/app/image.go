package app

import (
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"go-proxy-image-s3/internal/image"
	"go-proxy-image-s3/internal/s3"

	"github.com/go-chi/chi/v5"
)

func (app *App) ImageHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	key := r.URL.Query().Get("key")
	resolution := chi.URLParam(r, "resolution")

	extension := strings.ToLower(filepath.Ext(key))

	if !validGeometry(resolution) {
		geometryError := errors.New("invalid geometry")
		app.ServerErrorResponse(w, r, geometryError)
		return
	}

	if !validExtension(extension) {
		extensionError := errors.New("invalid extension")
		app.ServerErrorResponse(w, r, extensionError)
		return
	}

	// download image from s3
	image_file, err := s3.DownloadImage(key, app.Logger)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}

	// trasnform image
	transformedImage, err := image.Transform(*image_file, resolution)
	if err != nil {
		app.ServerErrorResponse(w, r, err)
		return
	}

	// return image
	http.ServeFile(w, r, *transformedImage)

	os.Remove(*image_file)
	os.Remove(*transformedImage)

	app.Logger.Infof("%v %s %s %s", time.Since(start), r.Method, r.URL.Path, key)
}

func validGeometry(geometry string) bool {
	geometryRE, _ := regexp.Compile(`(\d+)x(\d+)|original`)
	result := geometryRE.FindStringSubmatch(geometry)

	return len(result) > 0
}

func validExtension(extension string) bool {
	extensionRE, _ := regexp.Compile(`jpg|jpeg|png|gif`) // Prepare our regex
	result := extensionRE.FindStringSubmatch(extension)

	return len(result) > 0
}
