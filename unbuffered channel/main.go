package main 

import "fmt"
import "time"


func magic(ch chan int){

	fmt.Println("hello")

	ch <- 40 // block the code execution

	fmt.Println("ali")

}


func fun (ch chan int){

	fmt.Println("fun function")

	v := <- ch

	fmt.Println("the value of v : ",v)

}

func main(){

	var ch = make(chan int)

	fmt.Println("Start main function")

	go magic(ch)
	go fun(ch)


	fmt.Println("end-1")
	fmt.Println("end-2")
	fmt.Println("end-3")
	fmt.Println("end-4")

	time.Sleep(3*time.Second)

}