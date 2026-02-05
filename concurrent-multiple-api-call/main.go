package main

import "fmt"
import "io"
import "sync"
import "net/http"
import "encoding/json"

type PostIdStruct struct{
	 ID int `json:"id"`
}


type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}



func getSingleData(id int, wg *sync.WaitGroup, ch chan Todo) {
    defer wg.Done()

    URL := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", id)

    res, err := http.Get(URL)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer res.Body.Close()

    jsonDataSingle, err := io.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
        return
    }

    var todo Todo
    err = json.Unmarshal(jsonDataSingle, &todo)
    if err != nil {
        fmt.Println(err)
        return
    }

    ch <- todo
}



func main(){

	// declaration part 
	const URL string = "https://jsonplaceholder.typicode.com/todos"
	
	ch := make(chan Todo,500)
	var wg sync.WaitGroup

	// call the first api
	res,err := http.Get(URL)

	if(err != nil){
		fmt.Println("Find error on first api call : ",err)
		return
	}

	defer res.Body.Close()

	jsonData,err := io.ReadAll(res.Body)

	if(err != nil){
		fmt.Println("Find error to read the response data : ",err)
	}

	// fmt.Println(string(jsonData))

	var postsIdArr [] PostIdStruct

    // JSON → Go struct array
    err1 := json.Unmarshal(jsonData, &postsIdArr)
    if err1 != nil {
        fmt.Println("Error:", err1)
        return
    }

	for _, post := range postsIdArr {

	// fmt.Println("ID:", post.ID)

		wg.Add(1)

		go getSingleData(post.ID,&wg,ch)

    }

	go func() {
		wg.Wait()
		close(ch)
	}()

	for data := range ch {
		fmt.Println(data) //read data from channel 
	}


}
