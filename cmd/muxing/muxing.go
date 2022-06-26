package main

import (
	"fmt"
	"io/ioutil"
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

	router.HandleFunc("/name/{PARAM}", returnParameter).Methods(http.MethodGet)
	router.HandleFunc("/bad", badRequest).Methods(http.MethodGet)

	router.HandleFunc("/data", returnFromRequestBody).Methods(http.MethodPost)
	router.HandleFunc("/headers", returnFromRequestHeaders).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
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

func returnParameter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	finalString := "Hello, " + params["PARAM"] + "!"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(finalString))
}

func badRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func returnFromRequestBody(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	finalString := "I got message:\n" + string(body)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(finalString))
}

func returnFromRequestHeaders(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.Atoi(r.Header.Get("a"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	b, err := strconv.Atoi(r.Header.Get("b"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result := a + b

	w.Header().Add("a+b", strconv.Itoa(result))
	w.WriteHeader(http.StatusOK)
}
