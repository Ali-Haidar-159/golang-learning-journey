package main 

import "fmt"

type Shape interface{
	Area() float64
}

type Square struct{
	//Area
	// eta Shape type er hye gelo karon Area k attach kora hyse
	length float64
}

type Circle struct{
	radious float64
}

func (sqr Square) Area () float64{ // 1.Square struct er sathe attach korlm, jehetu area interface er member tai structure ta interface type mane Shape type er hye gelo
	fmt.Println("The area of a square : ",sqr.length * sqr.length)

	return sqr.length * sqr.length
}

func (c Circle) Area() float64{
	fmt.Println("The area of a circle : ",c.radious*c.radious*3.1416)

	return c.radious*c.radious*3.1416
}

func calculateArea (shp Shape){
	shp.Area()
}

func main(){

	var s1 Square = Square{length:4}
	var c1 Circle = Circle{radious:4}

	calculateArea(s1)
	calculateArea(c1)

}

