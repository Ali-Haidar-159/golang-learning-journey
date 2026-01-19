package main 

import "fmt"
import "ali.com/functions/service"

func main(){

	resu := service.Add(5,2)

	fmt.Println(resu)

}