package main

import "fmt"
import "time"

func magic (){

	for i:=0;i<100;i++{
		fmt.Println(i," : ","Hello")
	}

}

func fun (){

	for i:=0;i<100;i++{
		fmt.Println(i," : ","ali")
	}

}

func main(){

	go magic()
	go fun()

	time.Sleep(1 * time.Second) // wait to end the goroutine 

	fmt.Println("End of the main function")

}



