package main

import "fmt"
import "time"


func task1(ch chan string){

	time.Sleep(2 * time.Second)

	ch <- "task_1 is completed!"

	// close(ch)

}

func task2(ch chan string){
	time.Sleep(3 * time.Second)

	ch <- "task_2 is completed!"

	// close(ch)

}


func main(){

	ch1:= make(chan string)
	ch2:= make(chan string)

	go task1(ch1)
	go task2(ch2)


	for{
		select{
		case msg,ok:= <-ch1 :
			if !ok{
				ch1 = nil
			} else {
				fmt.Println("Task 1 is completed : ",msg)
				close(ch1)
			}

		case msg,ok := <-ch2 : 
			if (!ok){
				ch2 = nil
			} else {
				fmt.Println("Task 2 is completed : ",msg)
				close(ch2)
			}	
		}

		if(ch1==nil && ch2 == nil){

			fmt.Println("All task is completed")

			break 
		}

	}
}