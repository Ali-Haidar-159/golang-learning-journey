package main 

import "fmt"

func divideTwoNumbers (a int , b int)(int,error){

	if(b == 0){
		return 0 , fmt.Errorf("Something wrong , the value of b is : %d",b)
	} else{
		return a/b , nil
	}

}

func main (){

	result,err := divideTwoNumbers(10,0)

	if(err != nil){
		fmt.Println("The error is : ",err)

		return

	}

	fmt.Println("The result is : ",result)

}