package main 

import "fmt"

func main(){

	fmt.Print("My name is Ali")
	fmt.Println(" Haidar.")

	// variable declare 

	var id int 
	var university_name string = "East West University"
	department := "Computer Science and Engineering"

	fmt.Printf("Enter your student id : ")
	fmt.Scanf("%d",&id)

	fmt.Print("\n=====================\n\n")
	fmt.Printf("University : %s \n Department : %s\nStudent ID : %d\n",university_name,department,id)
	fmt.Print("\n=====================\n\n")

	fullString := fmt.Sprintf("University : %s \n Department : %s\nStudent ID : %d\n",university_name,department,id)

	fmt.Print(fullString)

}