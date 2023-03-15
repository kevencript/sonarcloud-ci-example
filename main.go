package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Here we are just printing a random Hash in order to verify 
	// on Go tests if everything is working fine. This way, we can
	// implement this into the Sonarqube analysis (This is an EXAMPLE)
	http.HandleFunc("/first", handlerSecondHash)


	http.HandleFunc("/second", handlerFirstHash)

	http.ListenAndServe(":8080", nil)
}

func getFirstHash() string {
	return "8743b52063cd84097a65d1633f5c74f6"
}

func getSecondHash() string {
	return "0cc175b9c0f1b6a831c399e269772662"
}

func handlerFirstHash(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getFirstHash())
}

func handlerSecondHash(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, getSecondHash())
}

