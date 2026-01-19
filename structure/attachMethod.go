package main 

import "fmt"

type Car struct{

	Brand string
	Price int 

}

func (c Car) display (){

	fmt.Println(c.Brand)
	fmt.Println(c.Price)

}

func main(){

	var myCar Car
	myCar = Car{Brand:"Toyota",Price:2000000}
	myCar.display()

}