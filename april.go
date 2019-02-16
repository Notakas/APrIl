package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePageHandler)
	router.HandleFunc("/revert/{word}", wordRevertHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":2999", router))
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func wordRevertHandler(w http.ResponseWriter, r *http.Request) {
	word := mux.Vars(r)["word"]
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "If you revert "+word+" you get "+reverse(word))
}

func reverse(word string) string {
	r := []rune(word)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
