package main

import ( "testing" 
"strconv"
"strings"
"fmt")


//Set of Unit Tests for matrixService.go

func TestEchoFunction(t *testing.T) {	
	var records = [][]string{ {"1","2","3"}, {"4","5","6"}, {"7","8","9"}}
	var expected string
	for _, row := range records {
		expected = fmt.Sprintf("%s%s\n", expected, strings.Join(row, ","))
	}

	result := echo(records);

	if(result != expected){
		t.Errorf("echo was incorrect, got: %s, want: %s.", result, expected)
	}
	
}

func TestInvertFunction(t *testing.T) {
	var records = [][]string{ {"1","2","3"}, {"4","5","6"}, {"7","8","9"}}
	transposedRecords := [][]string{ {"1","4","7"}, {"2","5","8"}, {"3","6","9"}}	
	var expected string
	for _, row := range transposedRecords {
		expected = fmt.Sprintf("%s%s\n", expected, strings.Join(row, ","))
	}

	result := invert(records);

	if(result != expected){
		t.Errorf("invert was incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestFlattenFunction(t *testing.T) {
	var records = [][]string{ {"1","2","3"}, {"4","5","6"}, {"7","8","9"}}
	expected := "1,2,3,4,5,6,7,8,9"

	result := flatten(records)

	if(result != expected){
		t.Errorf("flatten was incorrect, got: %s, want: %s.", result, expected)
	}
}

func TestSumFunction(t *testing.T) {
	var records = [][]string{ {"1","2","3"}, {"4","5","6"}, {"7","8","9"}}
	expected := int64(45)

	result := sum(records)

	if(result != expected){
		t.Errorf("sum was incorrect, got: %s, expected: %s.", strconv.Itoa(int(result)), strconv.Itoa(int(expected)))
	}
}

func TestMultiplyFunction(t *testing.T) {
	var records = [][]string{ {"1","2","3"}, {"4","5","6"}, {"7","8","9"}}
	expected := int64(362880)

	result := multiply(records)

	if(result != expected){
		t.Errorf("multiply was incorrect, got: %s, expected: %s.", strconv.Itoa(int(result)), strconv.Itoa(int(expected)))
	}
}
