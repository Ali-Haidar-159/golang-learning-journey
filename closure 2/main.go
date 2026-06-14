package main

import "fmt"

const a int = 10

var p int = 100

func magic() func() {
	money := 100

	fun := func() {
		money = money + a + p
		fmt.Println("The value of money is : ", money)
	}

	return fun
}

func work() {
	increment1 := magic()
	increment1()
	increment1()

	increment2 := magic()
	increment2()
	increment2()

}

func main() {
	fmt.Println("This is main function")

	work()

}

func init() {
	fmt.Println("This is practice code of Closure and Heap")
}
