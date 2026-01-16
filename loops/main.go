package main 

import "fmt"

func main(){

	var num int
	sum := 0 //shortcut declaration and initialization

	fmt.Print("Enter a number to get sum : ")
	fmt.Scanf("%d",&num)

	for i:=0 ; i<=num ; i++ {

		sum = sum + i 

	}

	fmt.Println("Your sum is : ",sum)

}