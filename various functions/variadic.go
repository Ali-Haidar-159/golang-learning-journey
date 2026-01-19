// Variadic or variable length argument function
// the parameter acts like an  array

package main

import "fmt"


func sum (nums ... int) int{
	total := 0 

	for _,value := range nums{

		total += value

	}

	return total

}


func main(){

	var res int 
	res = sum(1,2,3,4,5)

	fmt.Println("The sum is : ",res)

}

