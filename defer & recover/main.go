package main 

import "fmt"

func main (){

	defer func (){
		err := recover()
		if(err != nil){
			fmt.Println("The error is : ",err)
		}
	}()


	fmt.Println("Start of the main function")

	panic("Something went wrong")

	fmt.Println("End of the main function")
}