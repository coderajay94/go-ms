package main

import (
	"encoding/json"
	"fmt"
	det "go-ms/details"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health...")
	response := map[string]string{
		"status":    "up",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the home page...")
	//w.WriteHeader(http.StatusOK)
	//might not show in network tab
	//so changing the status to failure
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "application is up and running.\n")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching details...")
	//one way to import
	//hostName, _ := details.GetHostName()

	//other way to import, without mentioned alias for importing package
	//direcly use package.function name
	hostName, _ := det.GetHostName()
	ip := det.GetIP()
	//fmt.Fprintf(w, "Host Name is:", hostName)
	//fmt.Fprintf(w, "IP Address is is:", det.GetIP())

	response := map[string]string{
		"hostName":  hostName,
		"ipAddress": ip.String(),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)
	log.Println("Server has started...")
	http.ListenAndServe(":80", r)
}
