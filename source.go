package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getJsonData(w http.ResponseWriter, r *http.Request) {
	//Header set to Json
	w.Header().Set("Content-Type", "application/json")
	keys, ok := r.URL.Query()["search"]

	if !ok || len(keys[0]) < 1 {
		log.Fatal("URL parameter is missing")
		return
	}

	//Parameter keys from URL
	searchQuery := string(keys[0])
	dbSearchResult := GetSearchResult(searchQuery)

	//Print json in web
	json.NewEncoder(w).Encode(dbSearchResult)
}
