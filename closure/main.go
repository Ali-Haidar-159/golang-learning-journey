package main 

import "fmt"
import "ali.com/closure/service"

func main(){

	bank := service.Bank(1000)

	fmt.Println(bank(50,true))
	fmt.Println(bank(100,true))
	fmt.Println(bank(250,false))

}