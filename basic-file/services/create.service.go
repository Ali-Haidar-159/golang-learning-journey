package services 

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
		
	"ali.com/file/models"
)

func WriteUserData(user models.User)bool{

	file,err := os.OpenFile(
		"user-data.txt" ,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND ,
		0644,
	) 

	if (err != nil){
		return false
	}

	defer file.Close()

	fmt.Fprintf(file,"Id : %s , Name : %s , Country : %s , IsCitizen : %t \n",user.Id,user.Name,user.Country,user.IsCitizen)

	return true 

}



func ReadUserAllData() []models.User {

	file, err := os.Open("user-data.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var users []models.User

	for scanner.Scan() {

		line := scanner.Text()

		parts := strings.Split(line, ",")

		var user models.User

		for _, part := range parts {

			part = strings.TrimSpace(part)

			keyValue := strings.Split(part, ":")

			if len(keyValue) != 2 {
				continue
			}

			key := strings.TrimSpace(keyValue[0])
			value := strings.TrimSpace(keyValue[1])

			switch key {
			case "Id":
				user.Id = value
			case "Name":
				user.Name = value
			case "Country":
				user.Country = value
			case "IsCitizen":
				boolVal, err := strconv.ParseBool(value)
				if err == nil {
					user.IsCitizen = boolVal
				}
			}
		}

		users = append(users, user)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Read error:", err)
	}

	return users
}