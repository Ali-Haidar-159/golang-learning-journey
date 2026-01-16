package main 

import "fmt"

func main(){

	var fruits [5] string = [5] string {"apple","mango","jack-fruit","banana","orange"}

	// print the array using for-range loop

	for _,value := range fruits{
		fmt.Println(value)
	}

}

