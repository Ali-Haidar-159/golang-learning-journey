package main 

import "fmt"

func main(){

	var mark float64  //main syntax to declare a variable
	isPass := false //shortcut to declare a variable 

	fmt.Print("Enter your mark : ") 
	fmt.Scanf("%f",&mark)

	if(mark >= 80){
		fmt.Println("Congratulations ! you got A+") 
		isPass = true 
	}else if(mark >=70 && mark<80){
		fmt.Println("You got A") 
		isPass = true
	}else{
		fmt.Println("You Passed")
		isPass = true
	}

	fmt.Println("Is Pass : ",isPass)

}