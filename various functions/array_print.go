package main 

import "fmt"

// print the array elements 
func displayArray (arr []int) bool {

	var length int = len(arr)

	if(length > 2){

		fmt.Println("The element of array is : ",arr)

		return true

	} else {

		return false

	}


}


func main(){

	var marks [] int
	var isPrinted bool 

	marks = [] int {10,20}
	isPrinted = displayArray(marks)

	if(isPrinted){
		fmt.Println("Array is printed")
	}else{
		fmt.Println("Array is very short.")
	}


}