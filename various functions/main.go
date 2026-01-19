package main 

import "fmt"

// user defined function
//named functions 

func addNumbers (a int , b int) int {

	var res int 
	res = a + b 

	return res

}


func main(){

	result := addNumbers(10,5)

	fmt.Println("The result is : ",result)

}


