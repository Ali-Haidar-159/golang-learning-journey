package main 

import "fmt"
import "errors"
import "example.com/bank/utils"

// global scop 

var c1 utils.Bank = utils.Bank{

	Account : "1023589" ,
	Balance : 1000,

}


// all functions

func deposite(amount int , customer *utils.Bank)(error,bool){

	if(amount < 500){

		err := fmt.Errorf("Amount should be greater than 500 tk")

		return err , false
	}

	customer.Balance  += amount

	return nil , true

}

func withdraw(amount int , customer *utils.Bank)(error,bool){

	if(amount < 500){

		err := fmt.Errorf("Amount should be greater than 500 tk")

		return err , false
	}

	customer.Balance  -= amount

	return nil , true

}

func operation(amount int,operation string)(error,bool){

	switch operation{

		case "deposite" :
			err,_ := deposite(amount,&c1)

			if(err != nil){
				err := fmt.Errorf("Find an error from deposite function : %w",err) //error wrapping 
				return err,false
			}

			return nil,true

		case "withdraw" :
			err,_ := withdraw(amount,&c1)

			if(err != nil){
				err := fmt.Errorf("Find an error from deposite function : %w",err) //error wrapping 
				return err,false
			}

			return nil,true

		default :
			err := fmt.Errorf("Give me correct command , deposite or withdraw")
			return err , false

	}

}

func main (){

	var amount int 
	var operationMethod string 

	fmt.Printf("Enter the amount to deposite : ")
	fmt.Scan(&amount)

	fmt.Printf("Enter the operation : ")
	fmt.Scan(&operationMethod)

	err,_ := operation(amount,operationMethod)

	if(err != nil){

		realError := errors.Unwrap(err)

		if(realError != nil){
			fmt.Println("Error : ",realError)
		} else {
			fmt.Println("Error : ",err)
		}

		
	} else {
		fmt.Println("TK deposited successfully")
		fmt.Println("Details : ",c1)

	}

}