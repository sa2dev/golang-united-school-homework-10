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

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/bad", handleBad).Methods(http.MethodGet)
	router.HandleFunc("/name/{PARAM}", handleNameParam).Methods(http.MethodGet)
	router.HandleFunc("/data", handleDataParam).Methods(http.MethodPost)
	router.HandleFunc("/headers", handleHeaders).Methods(http.MethodPost)
	router.NotFoundHandler = http.HandlerFunc(notFound)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func handleHeaders(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	a, _ := strconv.Atoi(header.Get("a"))
	b, _ := strconv.Atoi(header.Get("b"))
	w.Header().Add("a+b", strconv.Itoa(a+b))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func handleDataParam(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	fmt.Fprintf(w, "I got message:\n%s", body)
}

func handleNameParam(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	param := params["PARAM"]
	fmt.Fprintf(w, "Hello, %s!", param)
}

func handleBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
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
