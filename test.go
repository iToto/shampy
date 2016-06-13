package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/shampoo", ShampooIndex)
    router.HandleFunc("/shampoo/{id}", ShampooShow)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello there and welcome to Shampy!")
}

func ShampooIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "ShampooIndex")
}

func ShampooShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    shampooId := vars["id"]
	fmt.Fprintln(w, "Shampoo show:", shampooId)
}