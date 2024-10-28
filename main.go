package main

import "fmt"

/*

Build a calculator project that allows the user to enter  2 numbers and the operation wanted. The calculator should show 
the final result

*/

func performCalculations(n1 float32, n2 float32, o string) float32 {

var result float32


	if o == "+" {
		result = n1 + n2

	} else if o == "-" {
		result = n1 - n2
	
	} else if o == "*" {
		result = n1 * n2 

	} else if o == "/" {
		result = n1 / n2
	}  

	if o == "/" && n2 == 0 {
		fmt.Println("Error, you can not divide by zero")
	}

	return result 
	
		}
	




func main() {

var number1 float32
var number2 float32
var operation string
var result float32

fmt.Println("Please enter the first number")
fmt.Scan(&number1)

fmt.Println("Please enter the second number")
fmt.Scan(&number2)

fmt.Println("Please enter the operation")
fmt.Scan(&operation)

result =performCalculations(number1, number2, operation )

fmt.Println("The result is", result)

}