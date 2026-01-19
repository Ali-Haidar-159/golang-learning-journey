package main

import "fmt"

// Parent Struct
type Person struct {
    Name string
    Age  int
}

// Parent struct's method
func (p Person) Introduce() {
    fmt.Printf("My name is %s, Age %d \n", p.Name, p.Age)
}

// Child Struct 
type Employee struct {
    Person
    EmpID int
}

// Method Overriding in Employee
func (e Employee) Introduce() {
    fmt.Printf("My name is %s, Age %d and  ID %d\n", e.Name, e.Age, e.EmpID)
}

func main() {

    // Person object
    p := Person{Name: "Rahim", Age: 30}
    p.Introduce()

    // Employee object
    e := Employee{
        Person: Person{Name: "Karim", Age: 25},
        EmpID: 101,
    }

    // Overridden method call
    e.Introduce()

    // Parent method call 
    e.Person.Introduce()
}
