package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func echo(records [][]string) string {

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}

func invert(records [][]string) string {	

	transposedRecords := transpose(records)
	var response string
	for _, row := range transposedRecords {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}

func flatten(records [][]string) string {

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s,", response, strings.Join(row, ","))
	}

	response = strings.TrimSuffix(response, ",")

	return response
}

func sum(records [][]string) int64 {

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s,", response, strings.Join(row, ","))
	}
	response = strings.TrimSuffix(response, ",")

	var result []string
	result = strings.Split(response, ",")
	sum := int64(0)

	for i := range result {
		value, err := strconv.ParseInt(result[i], 10, 64)

		if err != nil {
			return 0
		}
		sum += value
	}

	return sum
}

func multiply(records [][]string) int64 {

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s,", response, strings.Join(row, ","))
	}

	response = strings.TrimSuffix(response, ",")

	var result []string
	result = strings.Split(response, ",")
	multiply := int64(1)

	for i := range result {
		value, err := strconv.ParseInt(result[i], 10, 64)

		if err != nil {
			return 0
		}
		multiply *= value
	}

	return multiply
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func extractTextFromFile(w http.ResponseWriter, r *http.Request) ([][]string, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return nil, errors.New(err.Error())
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return nil, errors.New(err.Error())
	}

	return records, nil
}

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}
