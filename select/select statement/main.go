package main 

// import all the packages 

import "fmt"
import "time"

// main codes here 

func task_1(ch1 chan string){

	time.Sleep(3 * time.Second)

	str := fmt.Sprintf("Task-1 is completed")

	ch1 <- str

}


func task_2(ch2 chan string){

	time.Sleep(6 * time.Second)

	str := fmt.Sprintf("Task-2 is completed")

	ch2 <- str

}


func main(){

	ch1 := make(chan string)
	ch2 := make(chan string)

	go task_1(ch1)
	go task_2(ch2)

	select {

	case msg:= <-ch1:
		fmt.Println(msg)
	case msg:= <-ch2 :
		fmt.Println(msg)

	}

	time.Sleep(5 * time.Second)

}



