package main

import (
	"net/http"
	"os"
	"path/filepath"
)

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Get absolute path
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		//If we failed get the absolute path respond with a 400 bad request and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//prepend the path with static directory
	path = filepath.Join(h.staticPath, path)

	//Check wether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		//if file does not exists, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		//if we got an error (means the file doesn't exists)
		// return 500 internal server error and stop.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Otherwise use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
