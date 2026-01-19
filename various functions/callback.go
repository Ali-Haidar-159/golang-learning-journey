package main 

import "fmt"

// callback functions 

func add (a int , b int) int {

	var res int 
	res = a + b 

	return res

}

func subtract (a int , b int) int{

	var res int 
	res = a - b

	return res

}

// higher order function 
func math (x int , y int , fun func(int,int)int)int{

	var result int 

	result = fun(x,y)

	return result

}


func main (){

	finalResultAdd := math(100,15,add)
	finalResultSub := math(100,15,subtract)

	// display all the result 
	fmt.Println("The result of addition is : ",finalResultAdd)
	fmt.Println("The result of subtraction is : ",finalResultSub)

}