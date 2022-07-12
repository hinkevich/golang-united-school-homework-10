package main

import (
	"fmt"
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
func handlerBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	w.Write([]byte("500"))
	return
}
func handlerName(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["PARAM"]
	w.WriteHeader(200)
	w.Write([]byte("Hello, " + name + "!"))

	return
}
func handlerEmpty(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)
	w.Write([]byte(""))

	return
}
func handlerDataPost(w http.ResponseWriter, r *http.Request) {
	data := mux.Vars(r)["PARAM"]

	w.WriteHeader(200)
	//date := r.FormValue("PARAM")
	//date, err := r.GetBody()
	//if err != nil {
	//	log.Fatal(err)
	//}

	w.Write([]byte("I got message:\n" + data))

	return
}

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/bad", handlerBad).Methods("GET")
	router.HandleFunc("/name/{PARAM}", handlerName).Methods("GET")
	router.HandleFunc("/", handlerEmpty).Methods("POST")
	router.HandleFunc("/data/{PARAM}", handlerDataPost).Methods("POST")
	router.HandleFunc("/", handlerBad).Methods("POST")
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
