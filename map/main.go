package main

import "fmt"

func main (){

	// create a map using normal way
	var color = map[string]int{
		"red" : 2,
		"tomato" : 3 ,
		"black" : 4 ,
		"white" : 1,
	}

	// create a map using short hand way 
	devices := map[string]string {

		"processor" : "core i5 gen7" ,
		"ram" : "16GB" ,
		"hard-disk" : "1TB" ,
		"monitor" : "28 inchi" ,
		"key-board" : "A4 Tech" ,
		"mouse" : "Fantech" ,

	}

	// fmt.Println("The value of color is : ",color)
	// fmt.Println("The value of device is : ",devices)

	for key,value :=range color{
		fmt.Println(key , ":" ,value)
	}

	for key,value :=range devices{
		fmt.Println(key , ":" ,value)
	}

}