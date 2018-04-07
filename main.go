package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

// this will eventually be an actual query (for example type = A stefanaleksic.com)
func getDNS(url string) ([]string, error) {
	return net.LookupHost(url)
}

func createMessage(body string) string {
	twimlResponse :=
		`<?xml version="1.0" encoding="UTF-8"?>
			<Response>
				<Message><Body>%s</Body></Message>
			</Response>	
		`
	return fmt.Sprintf(twimlResponse, body)
}

func dnsTwimlResponse(w http.ResponseWriter, r *http.Request) {
	messageBody := r.URL.Query().Get("Body")
	listOfResolutions, err := getDNS(messageBody)

	if err != nil {
		fmt.Fprintf(w, createMessage("An error has occured! Oh my god..."))
	}

	resultingString := ""
	for _, resolution := range listOfResolutions {
		resultingString += resolution + "\n"
	}

	fmt.Fprintf(w, createMessage(resultingString))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", dnsTwimlResponse).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
