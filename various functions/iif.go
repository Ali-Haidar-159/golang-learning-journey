package main 

import "fmt"

func main(){

	// Immediately Invokeable Function - IIF 

	// This function call automatically after defining the function

	func(id int , name string){

		fmt.Printf("My name is %s and id is %d.\n",name,id)

	}(101,"Ali")


}