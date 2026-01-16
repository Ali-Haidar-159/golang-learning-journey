package main 

import "fmt"
import "strings"

func main (){

	var str string = "Ali Haidar"

	var a string = "ali"
	var b string = "haidar"

	// basic operation with strings 
	fmt.Println("The length is : ",len(str))
	fmt.Println("The upper case string is : ",strings.ToUpper(str))
	fmt.Println("The value of 0th index is : ",str[0]) //this print the ASCII value of 0th index character
	fmt.Println("The value of 0th index is : ",string(str[0])) //this print the value of 0th index


	// compare to string 
	isSame := strings.Compare(a,b)
	if(isSame == 0)	{
		fmt.Println("a = b")
	}else{
		fmt.Println("a and b is different string")
	}

	// search a string 
	isContains := strings.Contains(str,"Ali")
	if(isContains)	{
		fmt.Println("Ali is present in the str string.")
	}else{
		fmt.Println("The word is not found !")
	}

	// split the words of a string 
	var words []string
	words = strings.Split(str," ")
	fmt.Println(words)

	// join all the element of an array 
	var wordCollection [] string = []string {"my","name","is","Ali"}
	var fullString = strings.Join(wordCollection," ") ;
	fmt.Println("The full string is : ",fullString)

}