# GoMatrix
It's a Golang API that process matrix from a given .csv file

## :hammer: Setup

```bash

# Install the dependencies
Make sure to have Go installed on your machine
You can download Go from Golang website at https://golang.org/dl/

# Clone the app
git clone https://github.com/HudsonLima/GoMatrix.git

# Run the app
cd .\GoMatrix\
go run .

# Run Unit Tests
cd .\GoMatrix\
go test
```

## URL

``` http://localhost:8080/ ```

### Endpoints:

| Method | Endpoint | Description |
| --- | --- | --- |
| GET | /api/echo | Returns the matrix as a string in matrix format |
| GET | /api/invert | Returns the matrix as a string in matrix format where the columns and rows are inverted|
| GET | /api/flatten | Returns the matrix as a 1 line string, with values separated by commas|
| GET | /api/sum | Returns the sum of the integers in the matrix |
| GET | /api/multiply | Returns the product of the integers in the matrix |



## See the Postman sample API calls here: 

## /api/echo

Returns the matrix as a string in matrix format

![image](https://user-images.githubusercontent.com/5431216/111933819-fda75580-8a9e-11eb-951c-3da7b8756950.png)


## /api/invert

Returns the matrix as a string in matrix format where the columns and rows are inverted

![image](https://user-images.githubusercontent.com/5431216/111933879-14e64300-8a9f-11eb-9373-97af2395cf02.png)

## /api/flatten

Returns the matrix as a 1 line string, with values separated by commas

![image](https://user-images.githubusercontent.com/5431216/111933685-b6b96000-8a9e-11eb-9777-6a1897034272.png)

## /api/sum

Returns the sum of the integers in the matrix

![image](https://user-images.githubusercontent.com/5431216/111933734-d2246b00-8a9e-11eb-9e66-21fa303f9133.png)

## /api/multiply

Returns the product of the integers in the matrix

![image](https://user-images.githubusercontent.com/5431216/111936789-2d595c00-8aa5-11eb-97d7-50858a14eb81.png)




## See the CURL command sample API calls here:

### http://localhost:8080/api/echo         
curl -F file=@"/Users/Hudson/Documents/GoLang/matrix.csv" "localhost:8080/api/echo"

### http://localhost:8080/api/invert       
curl -F file=@"/Users/Hudson/Documents/GoLang/matrix.csv" "localhost:8080/api/invert"

### http://localhost:8080/api/flatten      
curl -F file=@"/Users/Hudson/Documents/GoLang/matrix.csv" "localhost:8080/api/flatten"

### http://localhost:8080/api/sum          
curl -F file=@"/Users/Hudson/Documents/GoLang/matrix.csv" "localhost:8080/api/sum"

### http://localhost:8080/api/multiply 
curl -F file=@"/Users/Hudson/Documents/GoLang/matrix.csv" "localhost:8080/api/multiply"






