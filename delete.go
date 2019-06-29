package main

import (
	"log"
	"net/http"
	"os"
)

// it can't be named delete due to duplicate name
func deleteFile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if !check(path) {
		w.WriteHeader(404)
		_, _ = w.Write([]byte("404"))
		log.Printf("delete %s 404", path)
		return
	}
	err := os.RemoveAll("." + path)
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte("500"))
		log.Printf("delete %s 500", path)
	} else {
		_, _ = w.Write([]byte("success"))
		log.Printf("delete %s success", path)
	}
}
