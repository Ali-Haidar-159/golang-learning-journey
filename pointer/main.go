package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func displayObject(p *Person) {
	fmt.Println("The details of person : ", p)
	fmt.Println("The details of person : ", *p)
	fmt.Println("The name of person : ", (*p).Name)

	// change the value of name using pointer
	(*p).Name = "Haidar"

}

/*
func displayArray(arr *[5]int) {
	fmt.Println("The array is : ", arr)
	fmt.Println("The array is : ", *arr)
	fmt.Println("The array is : ", (*arr)[0])
}
*/

func main() {

	/*
		var x int = 10
		// var ptr1 *int = &x // declare and initialize a pointer
		ptr := new(int) // declare a pointer using new() function
		ptr = &x

		fmt.Println("===== using variable =====")
		fmt.Println("The value of X : ", x)
		fmt.Println("The address of X : ", &x)

		fmt.Println("===== using pointer =====")
		fmt.Println("The value of x : ", *ptr)
		fmt.Println("The address of x : ", ptr)
	*/

	/*
		var arr [5]int = [5]int{10, 20, 30, 40, 50}
		displayArray(&arr)
	*/

	var p1 Person = Person{Name: "Ali", Age: 25}
	displayObject(&p1)
	fmt.Println("The name after func call : ", p1.Name)

}
