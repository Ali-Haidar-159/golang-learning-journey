package main 

import "fmt"

func main(){

	var i int = 1

	for{ // infinite loop
		fmt.Println(i)

		if(i == 10){
			break // break statement 
		}

		i++

	}

}