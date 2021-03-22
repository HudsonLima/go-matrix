package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", homePageHandler)
	http.HandleFunc("/api/echo", echoHandler)
	http.HandleFunc("/api/invert", invertHandler)
	http.HandleFunc("/api/flatten", flattenHandler)
	http.HandleFunc("/api/sum", sumHandler)
	http.HandleFunc("/api/multiply", multiplyHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Fprintf(w, "\nThe avaialable api endpoints are")
	fmt.Fprintf(w, "\n/api/echo")
	fmt.Fprintf(w, "\n/api/invert")
	fmt.Fprintf(w, "\n/api/flatten")
	fmt.Fprintf(w, "\n/api/sum")
	fmt.Fprintf(w, "\n/api/multiply")
	fmt.Println("Endpoint Hit: homePage")
}
