package main 

import "fmt"

func main(){

	// create a slice
	var numbers []int = []int {10,20,30,40,50} //in a slice the size or length is not fixed .

	// direct print
	fmt.Println("The slice is : ",numbers)

	fmt.Printf("\n\n")

	// print the slice unsing for loop-1 
	var i int 
	for i=0 ; i<len(numbers) ; i++{
		fmt.Printf("%d\t",numbers[i])
	}

	fmt.Printf("\n\n")

	// print the slice using loop-2 (for-range)
	for _,value :=range numbers{
		fmt.Printf("%d\t",value)
	}

}