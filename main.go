package main

import (
	"encoding/json"
	"log"
	"net/http"

	// images "github.com/daniellauziere/go-server/internal/images"
	"github.com/daniellauziere/go-server/internal/images"
	"github.com/gorilla/mux"
)

// ImageAsset - image asset based on database construct
type ImageAsset struct {
	ImageURL string `json:"imageUrl"`
}

var err error

// CORS ERROR FIX

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "http://192.168.1.5:3000")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func cor(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	images.UploadFile(w, r)

}

func uploadBlogImage(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	var resp = images.UploadFile(w, r)
	// var response = images.AddImageToPost(w, r, db, resp)
	json.NewEncoder(w).Encode(resp)

}

func main() {

	// Init router
	r := mux.NewRouter()

	// Route handelers / Endpoints

	// For all preflight requests ( subrouter )
	s := r.Methods("OPTIONS").Subrouter()
	s.HandleFunc("/{one}", cor)
	s.HandleFunc("/{one}/{two}", cor)
	s.HandleFunc("/{one}/{two}/{three}", cor)

	r.HandleFunc("/uploadImage", uploadBlogImage).Methods("POST")

	// Static file server

	fs := http.FileServer(http.Dir("./static-files"))
	r.PathPrefix("/{1}/").Handler(fs)
	r.PathPrefix("/{1}/{2}/").Handler(fs)
	r.PathPrefix("/{1}/{2}/{3}/").Handler(fs)

	log.Fatal(http.ListenAndServe(":3627", r))

}
