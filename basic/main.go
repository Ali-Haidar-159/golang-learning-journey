package main 

import "fmt"

func main(){

	fmt.Println("=== Students Details ===")

	// declaration part 

	var university string = "East West University"
	var name string
	var id int64
	var isInput bool

	var mark [5]float64 = [5]float64 {50,45,46,47,48}
	 

	fmt.Print("Do you want to get input (1/0) : ")
	fmt.Scanln(&isInput)
	 
	// input section 

	if(isInput){
		fmt.Print("Enter your name : ")
		fmt.Scanln(&name)

		fmt.Print("Enter your id : ")
		fmt.Scanln(&id)

	}

	// output section 

	fmt.Println(university)
	fmt.Printf("My name is: %s, my ID is: %d, my mark is : %v\n", name, id, mark)

}