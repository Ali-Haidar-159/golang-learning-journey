package main 

import "fmt"

type Person struct{
	Name string 
	Age int
}

type Employee struct{
	Person
	Emp_id int
}


func main(){

	var e Employee
	e = Employee {
		Person:Person{Name:"Ali Haidar",Age:22},
		Emp_id:10101,
	}

	fmt.Println(e.Name)
	fmt.Println(e.Age)
	fmt.Println(e.Emp_id)

}