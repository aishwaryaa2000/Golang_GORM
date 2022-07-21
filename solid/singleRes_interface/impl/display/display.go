package display

import (
	"encoding/json"
	"fmt"
	// "log"
	// "os"
	// "strconv"
)

type Display struct{
	Ans interface{}
}

func (d Display) DisplayJson(){
	ansDisplay := d.Ans
    json_data, _ := json.Marshal(ansDisplay)
	fmt.Println("In json format : ",string(json_data))
}
func (d Display) DisplayTerminal(){
	fmt.Println("Inside terminal.The area is : ",d.Ans)

}

// func writeIntoFile(a interface{}){
// 	f, err := os.OpenFile("ans.txt",os.O_APPEND|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer f.Close()

// 	b:=a.(float64)
// 	ansDisp := strconv.FormatFloat(b,'f',2,32)
// 	ansDisp = "The answer is " + ansDisp + "\n"
// 	_, errr := f.WriteString(ansDisp)
// 	if errr != nil {
// 		log.Println(errr)
// 	}

// 	fmt.Println("Data written onto ans.txt")
	
// }