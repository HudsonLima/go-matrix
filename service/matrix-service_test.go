package service

import ( "testing" 
"strings"
"fmt")

var (
	matrixService = NewMatrixService()
	matrixChannel = make(chan string)
)


func TestEcho(t *testing.T) {	
	var records = [][]string{ {"1","2","3"}, {"4","5","6"}, {"7","8","9"}}
	var expected string
	for _, row := range records {
		expected = fmt.Sprintf("%s%s\n", expected, strings.Join(row, ","))
	}

	go matrixService.Echo(records, matrixChannel);	
	result := <-matrixChannel

	 if(result != expected){
		t.Errorf("invert was incorrect, got: %s, want: %s.", result, expected)
	 }	
}

func TestInvert(t *testing.T) {
	var records = [][]string{ {"1","2","3"}, {"4","5","6"}, {"7","8","9"}}
	transposedRecords := [][]string{ {"1","4","7"}, {"2","5","8"}, {"3","6","9"}}	
	var expected string
	for _, row := range transposedRecords {
		expected = fmt.Sprintf("%s%s\n", expected, strings.Join(row, ","))
	}

	go matrixService.Invert(records, matrixChannel);	
	result := <-matrixChannel

	if(result != expected){
		t.Errorf("invert was incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestFlatten(t *testing.T) {
	var records = [][]string{ {"1","2","3"}, {"4","5","6"}, {"7","8","9"}}
	expected := "1,2,3,4,5,6,7,8,9"

	go matrixService.Flatten(records, matrixChannel);	
	result := <-matrixChannel

	if(result != expected){
		t.Errorf("flatten was incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestSum(t *testing.T) {
	var records = [][]string{ {"1","2","3"}, {"4","5","6"}, {"7","8","9"}}
	expected := "45"

	go matrixService.Sum(records, matrixChannel);	
	result := <-matrixChannel

	if(result != expected){
		t.Errorf("flatten was incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestMultiply(t *testing.T) {
	var records = [][]string{ {"1","2","3"}, {"4","5","6"}, {"7","8","9"}}
	expected := "362880"

	go matrixService.Multiply(records, matrixChannel);
	result := <-matrixChannel

	if(result != expected){
		t.Errorf("flatten was incorrect, got: %s, want: %s.", result, expected)
	}
}
