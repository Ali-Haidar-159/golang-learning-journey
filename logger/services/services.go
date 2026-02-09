package services 

import "fmt"


func Add (number1 int , number2 int)(int,error){
	
	if(number1>0 && number2>0){

		return number1 + number2  , nil

	} else {
 
		err := fmt.Errorf("Invalid number.Both number should be greater than zero")

		return 0  , err

	}

}


func Multiply (number1 int , number2 int)(int,error){
	
	if(number1>0 && number2>0){

		return number1 * number2  , nil

	} else {
 
		err := fmt.Errorf("Invalid number.Both number should be greater than zero")

		return 0  , err

	}

}



