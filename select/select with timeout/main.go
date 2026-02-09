package main

import "fmt"
import "time"


func task1(ch chan string){

	time.Sleep(20 * time.Second)

	ch <- "task_1 is completed!"

	close(ch)

}

func task2(ch chan string){
	time.Sleep(30 * time.Second)

	ch <- "task_2 is completed!"

	close(ch)

}


func main(){

	ch1:= make(chan string)
	ch2:= make(chan string)

	go task1(ch1)
	go task2(ch2)

	select{
	case <-ch1 :
		fmt.Println("Task 1 is completed")
	case <-ch2 :
		fmt.Println("Task 2 is completed")
	case <- time.After(5 * time.Second) :
		fmt.Println("No channel is completed")
	}

	
}