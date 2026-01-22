package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Student struct {
	Name  string
	Age   int
	Email string
	City  string
	Zip   int
}

func main() {

	jsonString := `{
		"name": "Ali",
		"age": 22,
		"email": "ali@example.com",
		"address": {
			"city": "Dhaka",
			"zip": "1207"
		}
	}`

	tempMap := make(map[string]interface{})
	var std Student

	err := json.Unmarshal([]byte(jsonString), &tempMap)
	if err != nil {
		fmt.Println("find error : ", err)
		return
	}

	// Type assertion for top-level fields
	std.Name = tempMap["name"].(string)
	std.Age = int(tempMap["age"].(float64))
	std.Email = tempMap["email"].(string)

	// Nested map
	address := tempMap["address"].(map[string]interface{})
	std.City = address["city"].(string)

	// zip conversion from string to int
	zipStr := address["zip"].(string)
	std.Zip, _ = strconv.Atoi(zipStr)

	fmt.Printf("Student obj is : %+v\n", std)
}
