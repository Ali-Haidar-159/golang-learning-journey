package main 

// import all the packages 

import "fmt"
import "log"
import "os"
import "time"

import "ali.com/logger/services"

var t = time.Now()
var logFile *os.File


func init(){

	file,err := os.OpenFile("ali.log",os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)

	if (err != nil){
		fmt.Println("Find error to create the logger file : ",err)
		return 
	}

	log.SetOutput(file)

	logFile = file

}


func main(){ 

	var operation string 
	var number1 , number2 int

	defer logFile.Close()
	
	fmt.Print("Enter the type of operation , add || multi || both : ")
	fmt.Scan(&operation)

	if(operation != "add" && operation!= "multi" && operation!="both"){

		fmt.Println("Ivalid Command.Write add || multi || both")
		log.Println("Invalid Command Entered - ",operation)

		return
	}

	fmt.Print("Enter the first number : ") 
	fmt.Scan(&number1)

	fmt.Print("Enter the second number : ") 
	fmt.Scan(&number2)

	switch operation {

	case "add" :
		
		res,err := services.Add(number1,number2)

		if(err != nil){

			fmt.Println("Find error : ",err)
			log.Println("Find error to add operation.Error : ",err)

		} else {

			fmt.Println("The result is : ",res)
			log.Printf("Add two numbers %d + %d = %d",number1,number2,res)

		}

	case "multi" :
		
		res,err := services.Multiply(number1,number2)

		if(err != nil){

			fmt.Println("Find error : ",err)
			log.Println("Find error to Multiply operation.Error : ",err)

		} else {

			fmt.Println("The result is : ",res)
			log.Printf("Multiply two numbers %d * %d = %d",number1,number2,res)

		}
	
	case "both" :
		
		res1,err := services.Add(number1,number2)
		res2,err := services.Multiply(number1,number2)

		if(err != nil){

			fmt.Println("Find error : ",err)
			log.Println("Find error in operation.Error : ",err)

		} else {

			fmt.Println("The addition result is : ",res1)
			fmt.Println("The multiplication result is : ",res2)
			log.Printf("Add and multiply two numbers %d + %d = %d , %d * %d = %d",number1,number2,res1,number1,number2,res2)

		}
	}	

}










