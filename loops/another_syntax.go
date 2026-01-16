package main 

import "fmt"

func main(){

	// declaration part 

	var num int 
	sum := 0
	i := 0

	fmt.Print("Enter a number to get sum : ")
	fmt.Scanf("%d",&num)

	// main code 
	for i<=num{

		sum += i 
		i++

	}

	fmt.Println("The result is : ",sum)

}