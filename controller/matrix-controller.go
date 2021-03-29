package controller

import (
	"encoding/csv"
	"errors"	
	"fmt"
	"net/http"
	"github.com.com/HudsonLima/GoMatrix/service"
)

var (
	matrixService service.MatrixService
	matrixChannel = make(chan string)
    errorMessage = "It seems the file you uploaded is either empty or does not contain a matrix =/. Please upload a new one"
)

type controller struct{}

type MatrixController interface {
	GetMatrixEcho(response  http.ResponseWriter, request *http.Request) 
	GetMatrixInvert(response  http.ResponseWriter, request *http.Request) 
	GetMatrixFlatten(response  http.ResponseWriter, request *http.Request) 
	GetMatrixSum(response  http.ResponseWriter, request *http.Request) 
	GetMatrixMultiply(response  http.ResponseWriter, request *http.Request) 
}

func NewMatrixController(service service.MatrixService) MatrixController {
	matrixService = service
	return &controller{}
}


/*Returns the matrix as a string in matrix format.*/
func (*controller) GetMatrixEcho(response http.ResponseWriter, request *http.Request) {

	records, err := extractTextFromFile(response, request)

	if err != nil {
		response.Write([]byte(fmt.Sprintf("error %s\n ", err.Error())))
		return
	}	

	if(len(records) <= 1){
		response.Write([]byte(fmt.Sprintf(errorMessage)))
		return	
	}
	  
	go matrixService.Echo(records, matrixChannel)
	result := <-matrixChannel

	fmt.Fprint(response, result)
}

/*Returns the matrix as a string in matrix format where the columns and rows are inverted*/
func (*controller) GetMatrixInvert(response http.ResponseWriter, request *http.Request) {

	records, err := extractTextFromFile(response, request)

	if err != nil {
		response.Write([]byte(fmt.Sprintf("error %s\n ", err.Error())))
		return
	}
	
	if(len(records) <= 1){
		response.Write([]byte(fmt.Sprintf(errorMessage)))
		return	
	}

	go matrixService.Invert(records, matrixChannel)
	result := <-matrixChannel

	fmt.Fprint(response, result)
}

/*Returns the matrix as a 1 line string, with values separated by commas*/
func (*controller) GetMatrixFlatten(response http.ResponseWriter, request *http.Request) {

	records, err := extractTextFromFile(response, request)

	if err != nil {
		response.Write([]byte(fmt.Sprintf("error %s\n", err.Error())))
		return
	}

	if(len(records) <= 1){
		response.Write([]byte(fmt.Sprintf(errorMessage)))
		return	
	}

	go matrixService.Flatten(records, matrixChannel)
	result := <-matrixChannel
	
	fmt.Fprint(response, result)
}

/*Returns the sum of the integers in the matrix*/
func (*controller) GetMatrixSum(response http.ResponseWriter, request *http.Request) {

	records, err := extractTextFromFile(response, request)

	if err != nil {
		response.Write([]byte(fmt.Sprintf("error %s\n", err.Error())))
		return
	}

	if(len(records) <= 1){
		response.Write([]byte(fmt.Sprintf(errorMessage)))
		return	
	}

	go matrixService.Sum(records, matrixChannel)
	result := <-matrixChannel

	fmt.Fprint(response, result)
}

/*Returns the product of the integers in the matrix*/
func (*controller) GetMatrixMultiply(response http.ResponseWriter, request *http.Request) {

	records, err := extractTextFromFile(response, request)

	if err != nil {
		response.Write([]byte(fmt.Sprintf("error %s\n", err.Error())))
		return
	}

	if(len(records) <= 1){
		response.Write([]byte(fmt.Sprintf(errorMessage)))
		return	
	}

	go matrixService.Multiply(records, matrixChannel)
	result := <-matrixChannel
		
	fmt.Fprint(response, result)
}

func extractTextFromFile(response http.ResponseWriter, request *http.Request) ([][]string, error) {
	file, _, err := request.FormFile("file")
	if err != nil {
		response.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return nil, errors.New(err.Error())
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		response.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return nil, errors.New(err.Error())
	}

	return records, nil
}


