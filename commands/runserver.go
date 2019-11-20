package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
  "log"
  "encoding/json"
)

type Pokemon struct {
  Number  string `json:"id"`
  Name    string `json:"name"`
}

var pokemonList []Pokemon

func main() {
  r := mux.NewRouter()

  // endpionts
  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "API response!\nNow on multiple lines.")
  })
  r.HandleFunc("/pokemon", getAllPokemon).Methods("GET")
  r.HandleFunc("/pokemon/{id}", getPokemon).Methods("GET")
  r.HandleFunc("/pokemon/", addPokemon).Methods("POST")
  r.HandleFunc("/pokemon/{id}", updatePokemon).Methods("POST")
  r.HandleFunc("/pokemon/{id}", deletePokemon).Methods("DELETE")

  fmt.Println("Server listening!")
  log.Fatal(http.ListenAndServe(":5000", r))
}

func getAllPokemon(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(pokemonList)
}

func getPokemon(w http.ResponseWriter, r *http.Request) {
}

func addPokemon(w http.ResponseWriter, r *http.Request) {
}

func updatePokemon(w http.ResponseWriter, r *http.Request) {
}

func deletePokemon(w http.ResponseWriter, r *http.Request) {
}


