package main 

// import all the packages 
import "fmt"
import "sync"

// structures 

type NumberSet struct{
	num1 int 
	num2 int 
	res int 
}

// all the go routines 

func addNumber (numSet NumberSet,wg *sync.WaitGroup,ch chan NumberSet){

	defer wg.Done()

	numSet.res = numSet.num1 + numSet.num2 

	ch <- numSet

}

func main (){

	var wg sync.WaitGroup
	ch := make(chan NumberSet,3)

	for i:=0 ; i<3 ; i++{

		wg.Add(1)

		n1 := 5+i 
		n2 := 2+i 
		ns := NumberSet{n1,n2,0}

		go addNumber(ns,&wg,ch)

	}

	go func(){
		wg.Wait()
		close(ch)
	}()

	for data := range ch {
		fmt.Println(data)
	}

}


