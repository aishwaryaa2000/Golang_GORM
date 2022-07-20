package display

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Display struct{
	Ans float64
}
func DisplayAns(a float64,format string){

	if format=="json"{
	json := giveJson(a)
	fmt.Println("In json format : ",string(json))
	}else if format=="terminal"{
		fmt.Println("Inside terminal.The area is : ",a)
	}else if format=="file"{
		writeIntoFile(a)
	}

}

func giveJson(a float64) []byte{
	ansDisplay := Display{a}
	//convert struct into json object
    json_data, _ := json.Marshal(ansDisplay)
	return json_data
}

func writeIntoFile(a float64){
	f, err := os.OpenFile("ans.txt",os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
		str := "The area is : " + strconv.FormatFloat(a,'f',2,32)
        str = str + "\n"
		_, errr := f.WriteString(str)
		if errr != nil {
			log.Println(errr)
		}
	fmt.Println("Data written onto ans.txt")
	
}