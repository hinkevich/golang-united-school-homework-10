package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Trailer", "AtEnd1, AtEnd2")
	w.Header().Add("Trailer", "AtEnd3")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	w.WriteHeader(http.StatusInternalServerError)
	w.WriteHeader(200)
	w.Header().Set("AtEnd1", "value 1")
	io.WriteString(w, "Status:"+http.StatusText(http.StatusInternalServerError))
	w.Header().Set("AtEnd2", "value 2")
	w.Header().Set("AtEnd3", "value 3") // These will appear as trailers.
	return
}
func handlerTwo(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["PARAM"]
	w.WriteHeader(200)
	w.Write([]byte("body: Hello, " + name + "!"))

	return
}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/bad", handler).Methods("GET")
	router.HandleFunc("/name/{PARAM}", handlerTwo).Methods("GET")
	router.HandleFunc("/", handler).Methods("POST")
	router.HandleFunc("/", handler).Methods("POST")
	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
