package main

import "fmt"

type Person struct{
	Name string
	Age int 
	IsPassed bool
}

func displayData (p Person){

	fmt.Println("")
	fmt.Println("Name : ",p.Name)
	fmt.Println("Age : ",p.Age)
	fmt.Println("IsPassed : ",p.IsPassed)
	fmt.Println("")
}


func main (){

	// manually initialize 
	var p1 Person 
	p1.Name = "Ali"
	p1.Age = 20
	p1.IsPassed = true

	// initialize using literal 
	var p2 Person
	p2 = Person{Name:"Haidar",Age:22,IsPassed:false}

	// declare and initialize using short hand method 
	p3 := Person{Name:"Kamal",Age:27,IsPassed:true}

	fmt.Println("Person-1 :",p1)
	
	displayData(p2)
	displayData(p3)

}