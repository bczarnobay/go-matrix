# GoMatrix

In this project it is possible to manipulate an integer matrix to perform the sum of the values, obtain the product of them, invert, print and flat the values as well.

- [Installation and execution](#installation-and-execution)
- [API](#api)
  * [Requirements](#requirements)
  * [Curl](#curl)
  * [Tests](#tests)
- [Caveats](#caveats)

## Installation and execution

To setup the project all you need to do is clone this repository in your machine.

1. To execute the application, please make sure you have Go installed on your machine and that you have port 8080 available.
2. Go to the root folder of this repository on your machine and run 
```
go run .
```

## API

### Requirements
Given an uploaded csv file
```
1,2,3
4,5,6
7,8,9
```

1. Echo (given)
    - Return the matrix as a string in matrix format.
    
    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ``` 
2. Invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ``` 
3. Flatten
    - Return the matrix as a 1 line string, with values separated by commas.
    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ``` 
4. Sum
    - Return the sum of the integers in the matrix
    ```
    // Expected output
    45
    ``` 
5. Multiply
    - Return the product of the integers in the matrix
    ```
    // Expected output
    362880
    ``` 

### Curl

It is recommended to have Postman installed on your machine to make the requests.

#### POST /echo
```
curl --location --request POST 'localhost:8080/echo' \
--form 'file=@"<PATH_TO_YOUR_CSV_EXAMPLE>"'
```
#### POST /invert
```
curl --location --request POST 'localhost:8080/invert' \
--form 'file=@"<PATH_TO_YOUR_CSV_EXAMPLE>"'
```
#### POST /flatten
```
curl --location --request POST 'localhost:8080/flatten' \
--form 'file=@"<PATH_TO_YOUR_CSV_EXAMPLE>"'
```
#### POST /sum
```
curl --location --request POST 'localhost:8080/sum' \
--form 'file=@"<PATH_TO_YOUR_CSV_EXAMPLE>"'
```
#### POST /multiply
```
curl --location --request POST 'localhost:8080/multiply' \
--form 'file=@"<PATH_TO_YOUR_CSV_EXAMPLE>"'
```

### Tests
To execute the tests you can run the following command from root folder:
```
go test ./...
```

### Caveats
Some points were missing on this project implementation, these are some open points to be worked on:

#### Requirements
- validate if all characters are valid (integers)
- convert matrix to integer before performing operations
- convert output from integer back to string

#### Tests
- extend services test to more robust scenarios
- extend e2e tests to consider other scenarios rather than only input validations
