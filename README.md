## Flatenning arbitrary array in Golang
### Introduction
Hi all:), thanks for your time to review my solution. I wanted to write the first question in Golang to show you that I have experience with it. 

### Algorithm
It is very simple, most of the people would do it this way probably. There is a recursive function that Iterates through the given array and checking the types of them, if there is an array in the list it recalls the function. if the value is not array it adds it to result array.

### Installation
Make sure you have Go installed in your machine. If not then you can go https://go.dev/play/ here and paste my solution to run as well.
-To run the function with my input;
``` 
go run main.go

```
-To run the test for the function with my test cases;
```
go test -v ./...
```