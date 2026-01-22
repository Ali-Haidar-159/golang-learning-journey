package main 

import "fmt"
import "encoding/json"


type Subject struct{
	SubjectName string 
	Code int 
}

// type Studnet struct{

// 	Name string `json:"StudnetName"`
// 	Age int     `json:"Student_age"`
// 	Gmail string
// 	IsPassed bool 
// 	Subject
// }


type Studnet struct{

	Name string 
	Age int    
	Gmail string
	IsPassed bool 
	Subject
}

func (s Studnet) JOSNConverter ()([]byte,error){

	dataMap := map[string]interface{}{

		"StudentName" : s.Name,
		"Student_age" : s.Age,
		"Gmail-account" : s.Gmail,
		"IsPassed" : s.IsPassed ,
		"Course Name" : s.Subject.SubjectName ,
		"Course code" : s.Subject.Code,

	}

	return json.Marshal(dataMap)

}



func main(){

	var std Studnet 
	std = Studnet{	
		Name : "Ali Haidar",
		Age : 24,
		Gmail : "ali@gmail.com",
		IsPassed : true ,
		Subject : Subject{
			SubjectName : "Web Development",
			Code : 479,
		},
	}

	// fmt.Println(std)

	jsonData,err := std.JOSNConverter()

	if(err != nil){
		fmt.Println("The error is : ",err)
		return 
	}

	fmt.Println("Json converted data is = ",string(jsonData))

}