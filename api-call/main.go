package main 

import "fmt"
import "net/http"
import "io"

func forGetReq(url string) ([]byte,error){

	res,err := http.Get(url)

	defer res.Body.Close()

	if(err != nil){
		return nil , err
	}
	
	data,err := io.ReadAll(res.Body)

	if(err != nil){
		return nil , err
	}

	return data,nil

}

func main(){

	const URL string = "https://jsonplaceholder.typicode.com/posts"

	data,err := forGetReq(URL)

	if(err != nil){
		fmt.Println("The error is : ",err)
	}else{
		fmt.Println(string(data))
	}

	

}