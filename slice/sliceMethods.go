package main 

import "fmt"

func main(){

	// declare a slice in shortcut 
	students := []string {"ali","haidar"}

	var length , capacity int

	length = len(students)
	capacity = cap(students)

	fmt.Println("The length of slice is : ",length)
	fmt.Println("The capacity of slice is : ",capacity)
	
	students = append(students,"kamal")

	length = len(students)
	fmt.Println("The length of slice is : ",length)


}