package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Intro to Golang",
		Description: "Here is how it works",
	},
}

func homeLink(httpWriter http.ResponseWriter, httpRequest *http.Request) {
	fmt.Fprintf(httpWriter, "Welcome Home Nerd!")
}

func createEvent(httpWriter http.ResponseWriter, httpRequest *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		fmt.Fprintf(httpWriter, "Enter data the event title and description to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	httpWriter.WriteHeader(http.StatusCreated)

	json.NewEncoder(httpWriter).Encode(newEvent)
}

func getOneEvent(httpWriter http.ResponseWriter, httpRequest *http.Request) {
	eventID := mux.Vars(httpRequest)["id"]

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			json.NewEncoder(httpWriter).Encode(singleEvent)
		}
	}
}

func getAllEvents(httpWriter http.ResponseWriter, httpRequest *http.Request) {
	json.NewEncoder(httpWriter).Encode(events)
}

func main() {
	initEvents()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
