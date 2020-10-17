package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(httpWriter http.ResponseWriter, httpRequest *http.Request) {
	fmt.Fprintf(httpWriter, "Welcome Home Nerd!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}
