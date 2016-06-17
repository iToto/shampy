package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "strconv"

    "github.com/gorilla/mux"
)

var shampoos = map[int]Shampoo{1: {"someBrand"}}
 
type Shampoo struct {
    Brand string
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/shampoo", ShampooIndex)
    router.HandleFunc("/shampoo/{id}", GetShampoo).Methods("GET")
    router.HandleFunc("/shampoo/{id}", PostShampoo).Methods("POST")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello there and welcome to Shampy!")
}

func ShampooIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
    j, _ := json.Marshal(shampoos)
    w.Write(j)
}

func GetShampoo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    shampooId, _ := strconv.ParseInt(vars["id"], 10, 0)

	fmt.Fprintln(w, "Shampoo show:", shampoos[int(shampooId)])
}

func PostShampoo(w http.ResponseWriter, r *http.Request) {

}