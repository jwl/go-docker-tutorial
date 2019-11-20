package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
)

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "API response!\nNow on multiple lines.")
  })

  fmt.Println("Server listening!")
  http.ListenAndServe(":80", r)
}
