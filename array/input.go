package main 

import "fmt"

func main(){

	students := [3]string {"ali","kalam","jabbar"} // declare the array in shortcut way
	var marks [5]int 

	var i int 

	for i=0 ; i<len(marks) ; i++ {

		fmt.Printf("Enter %dth marks : ",i+1)
		fmt.Scanf("%d",&marks[i])

	}

	// print the array directly 
	fmt.Println("The name of students : ",students)

	// print all the marks 
	
	for index,value := range marks { // for-range loop

		fmt.Printf("%d) %d\n",index+1,value)

	}


}

