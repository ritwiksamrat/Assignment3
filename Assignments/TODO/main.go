package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ToDo struct {
	WorkName string `json:"wname"`
	Progress string `json:"prog"`
}

var tod []ToDo

func getTODOS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tod)
}

func getTODO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range tod {
		if item.WorkName == params["wname"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&ToDo{})
}

func createTODO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var to ToDo
	_ = json.NewDecoder(r.Body).Decode(&to)
	tod = append(tod, to)
	json.NewEncoder(w).Encode(to)
}

func updateTODO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range tod {
		if item.WorkName == params["wname"] {
			tod = append(tod[:idx], tod[idx+1:]...)
			var to ToDo
			_ = json.NewDecoder(r.Body).Decode(&to)
			to.WorkName = params["wname"]
			tod = append(tod, to)
			json.NewEncoder(w).Encode(to)
			return
		}
	}
}

func deleteTODO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range tod {
		if item.WorkName == params["wname"] {
			tod = append(tod[:idx], tod[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(tod)
}

func main() {
	r := mux.NewRouter()

	tod = append(tod, ToDo{WorkName: "Writing", Progress: "Started"})
	tod = append(tod, ToDo{WorkName: "Reading", Progress: "Finished"})
	tod = append(tod, ToDo{WorkName: "Development", Progress: "Proceeding"})

	r.HandleFunc("/todo", getTODOS).Methods("GET")
	r.HandleFunc("/todos/{wname}", getTODO).Methods("GET")
	r.HandleFunc("/todos", createTODO).Methods("POST")
	r.HandleFunc("/todos/{wname}", updateTODO).Methods("PUT")
	r.HandleFunc("/todos/{wname}", deleteTODO).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
