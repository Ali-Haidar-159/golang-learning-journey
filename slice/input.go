package main 

import "fmt"

func main(){

	// declare a slice in shortcut 
	students := []string {"ALI","HAIDAR"}

	// create another slice 
	var marks []int = make([]int , 5)

	// input in the slice 
	for i:=0 ; i<5 ; i++{
		fmt.Printf("Enter the %d value : ",i+1)
		fmt.Scanf("%d",&marks[i])
	}

	fmt.Println(students)
	fmt.Println(marks)

}


/* we can create a slice using 2 way 

1. directly --> 
var slice_name type [] = []type {}data
slice_name := []type {data}


2. using make() method -->
var slice_name []type = make ([]type,length,capacity)
slice_name = make ([]type,length,capacity)


*/



