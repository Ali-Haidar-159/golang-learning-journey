
package utils

// import all the packages

import "fmt"
import "strconv"
import "net/http"
import "io"
import "sync"
import "os"

import "ali.com/api_call/models"


func singleUserData(id string , wg *sync.WaitGroup){

	defer wg.Done() // execute at the end of function

	// Access environment variable
	api := os.Getenv("API")

	// var URL string = "https://jsonplaceholder.typicode.com/users/" + id 
	var URL string = api + id 

	res,err := http.Get(URL)

	if(err != nil){
		fmt.Println("Find error to get the single user data : ",err)
		return
	}

	resBody , err := io.ReadAll(res.Body)

	if(err != nil){
		fmt.Println("Find error to get the single user data : ",err)
		return
	}

	fmt.Println("=============================")
	fmt.Println(string(resBody))
	fmt.Println("")

}



func Display(user *[]models.User){ //user = pointer = 10011a10a

	// for _,usr :=range *user{

	// 	fmt.Println("Name : ",usr.Name)
	// 	fmt.Println("ID : ",usr.ID)

	// }

	var wg sync.WaitGroup

	wg.Add(10) //counter = 10

	for _,usr :=range *user{

		strId := strconv.Itoa(usr.ID)
		go singleUserData(strId,&wg)

	}

	wg.Wait()// when the value of counter=0 then the code executes 

}