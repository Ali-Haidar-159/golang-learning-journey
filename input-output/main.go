package main 

import "fmt"

func main(){

	var name string 
	var id int64
	var marks float64
	var isPass bool

	fmt.Print("Enter your name (one word) : ")
	fmt.Scanf("%s",&name) 

	fmt.Print("Enter your id : ")
	fmt.Scanf("%d",&id) 

	fmt.Print("Enter your marks : ")
	fmt.Scanf("%f",&marks)

	fmt.Print("Do you passed :")
	fmt.Scanf("%t", &isPass)

	//this is the shortcut to declare a variable
	fullString := fmt.Sprintf("My name is : %s , Id: %d , Marks : %f , IsPassed : %t\n",name,id,marks,isPass)

	fmt.Printf("%v",fullString)	

}