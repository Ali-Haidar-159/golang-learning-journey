package main 

import "fmt"

func magic (a int , b int , ch chan int){

	var res int = a + b 

	ch <- res

}

func main(){


	// this is buffered channel 
	chnl := make(chan int,5)

	go magic (2,3,chnl)
	go magic (2,4,chnl)
	go magic (2,5,chnl)
	go magic (2,6,chnl)
	go magic (2,7,chnl)

	

	for i:=0;i<5;i++{
		var res int = <-chnl
		fmt.Println("The result is : ",res)
		fmt.Println()
	}
	

}