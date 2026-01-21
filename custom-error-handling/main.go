package main 

import "fmt"
import "ali.com/custome_error/service"

func main (){

	// input the age via consol from the users 

	var age int 
	fmt.Print("Enter youe age to check the validity : ")
	fmt.Scanf("%d",&age)

	err := service.CheckAge(age)

	if(err != nil){
		fmt.Println("The error is : ",err)
		return
	}

	fmt.Println("Your age is valid , so you can proceed")

}