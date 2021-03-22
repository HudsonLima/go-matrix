package main

import (
	"fmt"
	"net/http"
)

const errorMessage = "It seems the file you uploaded is either empty or does not contain a matrix =/. Please upload a new one"

/*Returns the matrix as a string in matrix format.*/
func echoHandler(w http.ResponseWriter, r *http.Request) {

	records, err := extractTextFromFile(w, r)

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s\n ", err.Error())))
		return
	}	

	if(len(records) <= 1){
		w.Write([]byte(fmt.Sprintf(errorMessage)))
		return	
	}

	response := echo(records)
	fmt.Fprint(w, response)
}

/*Returns the matrix as a string in matrix format where the columns and rows are inverted*/
func invertHandler(w http.ResponseWriter, r *http.Request) {

	records, err := extractTextFromFile(w, r)

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s\n ", err.Error())))
		return
	}
	
	if(len(records) <= 1){
		w.Write([]byte(fmt.Sprintf(errorMessage)))
		return	
	}

	response := invert(records)
	fmt.Fprint(w, response)
}

/*Returns the matrix as a 1 line string, with values separated by commas*/
func flattenHandler(w http.ResponseWriter, r *http.Request) {

	records, err := extractTextFromFile(w, r)

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s\n", err.Error())))
		return
	}

	if(len(records) <= 1){
		w.Write([]byte(fmt.Sprintf(errorMessage)))
		return	
	}

	response := flatten(records)
	fmt.Fprint(w, response)
}

/*Returns the sum of the integers in the matrix*/
func sumHandler(w http.ResponseWriter, r *http.Request) {

	records, err := extractTextFromFile(w, r)

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s\n", err.Error())))
		return
	}

	if(len(records) <= 1){
		w.Write([]byte(fmt.Sprintf(errorMessage)))
		return	
	}

	response := sum(records)
	fmt.Fprint(w, response)
}

/*Returns the product of the integers in the matrix*/
func multiplyHandler(w http.ResponseWriter, r *http.Request) {

	records, err := extractTextFromFile(w, r)

	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s\n", err.Error())))
		return
	}

	if(len(records) <= 1){
		w.Write([]byte(fmt.Sprintf(errorMessage)))
		return	
	}

	response := multiply(records)
	fmt.Fprint(w, response)
}


