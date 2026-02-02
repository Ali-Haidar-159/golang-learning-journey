package main 

import "fmt"

func fun (x * int){
	*x = 100
}

func main (){

	var a int = 10 

	fmt.Println("The value of a : ",a)
	fmt.Println("The address of a : ",&a)

	fun(&a)

	fmt.Println("============ After calling the function ===========")

	fmt.Println("The value of a : ",a)
	fmt.Println("The address of a : ",&a)

}