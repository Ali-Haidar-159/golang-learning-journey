package main 

import "fmt"

func main(){

	// declare and initialize a map using normal way 
	var students map[string]map[string]int
	students = map[string]map[string]int{

		"Ali" : {
			"math" : 80 ,
			"english" : 87 ,
			"cse" : 90 ,
		} ,
		"Haidar" : {
			"math" : 84 ,
			"english" : 85 ,
			"cse" : 89,
		} ,

	}

	// declare a map using make function
	var teacherNum map[string]int 
	teacherNum = make(map[string]int)

	// initialize that map 
	teacherNum["math"] = 5
	teacherNum["english"] = 4 
	teacherNum["cse"] = 6

	// print all the map 
	fmt.Println("The detail of Ali : ",students["Ali"])
	fmt.Println("The detail of Haidar : ",students["Haidar"])

	fmt.Println("") //for new line 

	fmt.Println("The number of teacher in each department :",teacherNum)


}