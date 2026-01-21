package service 

import "fmt"

// create a custom error , if age<18 then this error type works 

type AgeError struct{
	Age int 
}

func (ae AgeError) Error () string { //implement the Error method of error interface
	
	errMessage := fmt.Sprintf("Age = %d , age should be greater than 18 ." , ae.Age)

	return errMessage

}

