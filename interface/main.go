package main 

import "fmt"

type Shape interface{
	Area() float64
}

type Rectangle struct{
	length float64
	width float64
}

func (r Rectangle) Area() float64{

	return r.length * r.width

}

func main (){
	rec := Rectangle{length:10,width:4}

	var area float64 = rec.Area()

	fmt.Println("The area of rectangle is : ",area)
}