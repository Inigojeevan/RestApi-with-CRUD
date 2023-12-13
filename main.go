package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var profiles []Profile

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type Profile struct {
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Employee    User   `json:"employee"`
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var newProfile Profile
	json.NewDecoder(r.Body).Decode(&newProfile)

	w.Header().Set("Content-Type", "application/json")

	profiles = append(profiles, newProfile)

	json.NewEncoder(w).Encode(profiles)
}

func getAllProfiles(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(profiles)
	w.Header().Set("Content-Type", "application/json")
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	var idParm string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParm)
	if err != nil {
		w.Write([]byte("id not created"))
		w.WriteHeader(400)
		return
	}

	if id >= len(profiles) {
		w.Write([]byte("invalid ID"))
		w.WriteHeader(404)
		return
	}

	profile := profiles[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	var idParm string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParm)

	if err != nil {
		w.Write([]byte("ID not created"))
		w.WriteHeader(400)
		return
	}

	if id >= len(profiles) {
		w.Write([]byte("Invalid ID"))
		w.WriteHeader(400)
		return
	}

	var updateProfile Profile
	json.NewDecoder(r.Body).Decode(&updateProfile)

	profiles[id] = updateProfile

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateProfile)

}

func deleteProfile(w http.ResponseWriter, r *http.Request) {
	var idParm string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParm)

	if err != nil {
		w.Write([]byte("ID not created"))
		w.WriteHeader(400)
		return
	}

	if id >= len(profiles) {
		w.Write([]byte("Invalid ID"))
		w.WriteHeader(400)
		return
	}
	profiles = append(profiles[:id], profiles[id+1:]...)
	w.WriteHeader(200)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/profiles", addItem).Methods("POST")
	router.HandleFunc("/profiles", getAllProfiles).Methods("GET")
	router.HandleFunc("/profiles/{id}", getProfile).Methods("GET")
	router.HandleFunc("/profiles/{id}", updateProfile).Methods("PUT")
	router.HandleFunc("/profiles/{id}", deleteProfile).Methods("DELETE")
	http.ListenAndServe(":5000", router)
}
