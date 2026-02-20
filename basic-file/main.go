package main

import (
	"fmt"
	"os"

	"github.com/google/uuid"

	"ali.com/file/models"
	"ali.com/file/services"
)

func main(){

	var userInfo models.User
	var target string
	
	fmt.Printf("Enter your target (add/read): ")
	fmt.Scan(&target)

	switch target{

	// switch os.Args[1] {

	case "add" :
		userInfo = getDataFromUser()
		isSucessfullyWrite := services.WriteUserData(userInfo)

		if(isSucessfullyWrite){
			fmt.Println("User data stored in Database.")
		}

	case "read" :
		allUsers := services.ReadUserAllData()

		for _, user := range allUsers {
			fmt.Printf("%+v\n", user)
		}

	default:
		fmt.Println("Invalid target entered !!!")

	}


}



func getDataFromUser()models.User{

	var isCitizen int
	var user models.User

	user.Id = uuid.New().String()

	fmt.Printf("Enter your name : ")
	fmt.Scan(&user.Name)

	fmt.Printf("Enter your country : ")
	fmt.Scan(&user.Country)

	fmt.Printf("Are you citizen of this country (1/0) :")
	fmt.Scan(&isCitizen)

	if (isCitizen == 1){
		user.IsCitizen = true
	} else {
		user.IsCitizen = false
	}

	return user

}







