package main 

import "fmt"

// anonymous function expression

var fun = func (str string) {
	fmt.Println("The String is : ",str)
}

func main(){

	fun("My name is Ali Haidar")

}