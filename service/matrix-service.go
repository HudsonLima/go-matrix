package service

import (
	"fmt"
	"strconv"
	"strings"
)

type service struct{}

type MatrixService interface {
	Echo(records [][]string, channel chan string)  
	Invert(records [][]string, channel chan string)
	Flatten(records [][]string, channel chan string)
	Sum(records [][]string, channel chan string) 
	Multiply(records [][]string, channel chan string) 
}

func NewMatrixService() MatrixService {
	return &service{}
}


func (*service) Echo(records [][]string, channel chan string) {

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	channel <- response
}

func (*service) Invert(records [][]string, channel chan string) {	

	transposedRecords := Transpose(records)
	var response string
	for _, row := range transposedRecords {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}

	channel <- response
}

func (*service) Flatten(records [][]string, channel chan string) {

	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s,", response, strings.Join(row, ","))
	}
	
	channel <- strings.TrimSuffix(response, ",")	
}

func (*service) Sum(records [][]string, channel chan string)  {

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
			return 
		}
		sum += value
	}

	channel <- strconv.FormatInt(sum, 10)
}

func (*service) Multiply(records [][]string, channel chan string)  {

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
			return 
		}
		multiply *= value
	}

	channel <- strconv.FormatInt(multiply, 10)
}

func Transpose(slice [][]string) [][]string {
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
