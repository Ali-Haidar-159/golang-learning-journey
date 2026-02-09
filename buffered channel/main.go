package main 

import "fmt"
import "time"

func magic (ch chan int){
	for i:=1 ; i<=4 ; i++ {
		fmt.Println("Sending - ",i)
		ch <- i
		fmt.Println("Sent - ",i)
	}
}
func fun (ch chan int){
	for i:=1 ; i<=2 ; i++{
		v := <-ch
		fmt.Println("the value of v : ",v)
	}
}
func main (){
	ch := make(chan int,2)
	go magic(ch)
	go fun(ch)
	time.Sleep(4 * time.Second)

}