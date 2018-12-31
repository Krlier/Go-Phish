package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	respond "gopkg.in/matryer/respond.v1"
)

// GetSite is responsible for handling the request for the Levenshtein distance.
func GetSite(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Levenshtein!")
	vars := mux.Vars(r)
	siteName := vars["siteName"]

	fmt.Println("Website For Testing:", siteName)

	sResultPercentage, sResultName := GoLevenshtein(siteName)
	s := fmt.Sprintf("The website with the smallest difference is: %s, being %.2f%% different", sResultName, sResultPercentage)
	respond.With(w, r, http.StatusOK, s)

	return
}

// handler is the healthcheck route.
func handler(w http.ResponseWriter, r *http.Request) {
	respond.With(w, r, http.StatusOK, "WORKING!")
	fmt.Print("WORKING!")
	return
}

// endHandler ends the API gracefully.
func endHandler(w http.ResponseWriter, r *http.Request) {
	respond.With(w, r, http.StatusOK, "Ending Gracefully!")
	fmt.Println("Ending Gracefully!")

	os.Exit(0)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/leven/{siteName}", GetSite)
	r.HandleFunc("/end", endHandler)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8888", r))
}
