package main 

import "fmt"

func main(){

	var fruits [5] string = [5] string {"apple","mango","jack-fruit","banana","orange"}

	// direct print 
	fmt.Println("The fruit list is : ",fruits)

	// print the array using for loop

	length := len(fruits)

	for i:=0 ; i<length ; i++{
		fmt.Printf("%s\t",fruits[i])
	} 

}

