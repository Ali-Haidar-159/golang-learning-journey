package main 

import "fmt"
import "errors"

func calculateNumber (num1 int , num2 int) (int,error) {

	if(num2 == 0){

		err := fmt.Errorf("find error because second number is ZERO.")

		return 0,err
	}

	return num1/num2 , nil

}

func loadApp (num1 int , num2 int) (int,error){

	result,err := calculateNumber(num1,num2)

	if(err != nil){
		
		errW := fmt.Errorf("In calculate number function return an error : %w",err)

		return 0 , errW
	}

	return result , nil

}

func main (){


	var num1 , num2 int 

	fmt.Print("Enter the first number : ")
	fmt.Scanf("%d",&num1)

	fmt.Print("Enter the second number : ")
	fmt.Scanf("%d",&num2)

	result,err := loadApp(num1,num2)

	if(err != nil){
		
		fmt.Println("The wrapped error is : ",err)

		// unwrapping the error 

		mainError := errors.Unwrap(err)
		fmt.Println("Main error is : ",mainError)

		return 
		
	}

	fmt.Println("The result is : ",result)

}