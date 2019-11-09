package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func main() {
	router := mux.NewRouter()

	//URL Routes
	router.HandleFunc(searchRouteUrl, getJsonData).Methods("GET")

	spa := spaHandler{staticPath: staticPathDir, indexPath: indexPathFile}
	router.PathPrefix("/").Handler(spa)

	//Server conf
	srv := &http.Server{
		Handler:      router,
		Addr:         serverAddress,
		WriteTimeout: timeout,
		ReadTimeout:  timeout,
	}

	log.Fatal(srv.ListenAndServe())
}
