package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pradumnasaraf/go-api/helper"
	"github.com/pradumnasaraf/go-api/model"
)

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	helper.InsertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	helper.UpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	helper.DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMyAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	deltedCount := helper.DeleteAllMovies()
	json.NewEncoder(w).Encode(deltedCount)
}

func GetMyAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := helper.GetAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func GetAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	movie := helper.GetOneMovie(params["id"])
	json.NewEncoder(w).Encode(movie)
}

func ServeHomepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is working fine 🚀."))
}
