package main 

import "fmt"
import "time"

func main(){

	ticker := time.NewTicker(1 * time.Second)

	// count := 0

	targetHour := 15
	targerMin := 37
	targetSec := 12

	

	for t := range ticker.C{
		fmt.Println(t)
		// count++ 

		// if(count == 5){
		// 	break
		// }

		if(t.Hour()== targetHour && t.Minute()==targerMin && t.Second()==targetSec){
			fmt.Println("Alarm is ringing")
			break
		}
	}

}