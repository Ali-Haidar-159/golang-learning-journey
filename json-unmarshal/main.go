package main 

import "fmt"
import "encoding/json"


type Student struct{

	Name string `json:"full_name"`
	Age int `json:"age_"`

}


func main(){

	jsonData := `{"full_name":"Ali Haidar","age_":23}`

	var std Student

	err:= json.Unmarshal([]byte(jsonData),&std)

	if(err != nil){
		fmt.Println("find error : ",err)
	}else{
		fmt.Println(std)
	}



}