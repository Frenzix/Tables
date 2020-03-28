package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
)

type event struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	ID          string `json:"ID"`
}

type AllEvents struct{
	Events []event
}

var events = AllEvents{
        Events: []event{ {
            ID:          "1",
            Title:       "Read Below!",
            Description: "Learn how to change data to this table from the information below!",
        },
    },
}


func homeLink(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("application_botsp.html")

	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	err1 := t.Execute(w,events)
	
	if err1 != nil{
	 
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalid data")
	}
	
	json.Unmarshal(reqBody, &newEvent)
	events.Events = append(events.Events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func getAllEvents(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)
}

func getOneEvent(w http.ResponseWriter, r *http.Request) {

	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events.Events {
		if singleEvent.ID == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	var updatedEvent event
	eventID := mux.Vars(r)["id"]

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Invalid Data")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range events.Events {
		if singleEvent.ID == eventID {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			events.Events = append(events.Events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for i, singleEvent := range events.Events {
		if singleEvent.ID == eventID {
			events.Events = append(events.Events[:i], events.Events[i+1:]...)
			fmt.Fprintf(w, "The event with ID %v has been deleted successfully", eventID)
		}
	}
}


func main() {
	port := 8080
	router := mux.NewRouter().StrictSlash(true)
	fmt.Printf("Running on port %d", port)
	router.HandleFunc("/", homeLink).Methods("GET")
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events", getAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", updateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), router))
}

//curl -X PATCH -H 'Content-type:application/json' -d '{"ID":"IDHERE", "Title":"TITLEHERE", "Description":"DESHERE"}' http://127.0.0.1:8000/{id}
//curl -X POST -H 'Content-type:application/json' -d '{"ID":"IDHERE", "Title":"TITLEHERE", "Description":"DESHERE"}' http://localhost:8000/event