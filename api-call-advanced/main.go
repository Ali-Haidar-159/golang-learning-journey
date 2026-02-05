package main

// import all the packages and modules 
import "fmt"
import "net/http"
import "io"
import "encoding/json"
import "github.com/joho/godotenv"

import "ali.com/api_call/models"
import "ali.com/api_call/utils"

// others code 



// main function 

const URL = "https://jsonplaceholder.typicode.com/users"

func init(){

	err := godotenv.Load()

	if err != nil {
		fmt.Println("find error to load the env file : ",err)
		return
	}

}


func main(){

	var user []models.User

	res,err := http.Get(URL)

	if(err != nil){
		fmt.Println("Find error to get data : ",err)
		return 
	}  

	resBody,err := io.ReadAll(res.Body)

	if(err != nil){
		fmt.Println("Find error to read the response body : ",err)
		return 
	}

	json.Unmarshal(resBody,&user)

	utils.Display(&user)

}




