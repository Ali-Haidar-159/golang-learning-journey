package main 

// import all the module or packages

import "fmt"
import "sync"

// all the go routines 

func task(id int , wg *sync.WaitGroup){

	defer wg.Done()

	fmt.Println("Task - ",id," is completed !")

}

// main code 

func main(){

	var wg sync.WaitGroup

	wg.Add(3)

	task(1,&wg)
	task(2,&wg)
	task(3,&wg)

	go func(){  //IIF
		wg.Wait()
	}()

	fmt.Println("End of main function")
}

