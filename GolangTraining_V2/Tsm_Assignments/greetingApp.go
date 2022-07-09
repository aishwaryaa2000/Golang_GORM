package main

import( 
		"fmt"
		"time"
)

func greeting() string{

	currentHour:= time.Now().Hour()
  	currentSecond := time.Now().Second()

	var greetingMsg string

	if (currentHour==6 && currentSecond!=0)||(currentHour==11 && currentSecond==0)||(currentHour > 6 && currentHour<11){
		// 6:00:01 am to 11:00:00 am 
		greetingMsg = "Good morning"
	} else if (currentHour >=11 && currentHour<16) || (currentHour==16 && currentSecond==0){
		//11:00:01 am to 4:00:00 pm 
		greetingMsg =  "Good Afternoon"
	} else if (currentHour >= 16 && currentHour<21) || (currentHour==21 && currentSecond==0){
		//4:00:01 pm to 9:00:00 pm
		greetingMsg = "Good Evening"
	} else {
		//9:00:01 pm to 6:00:00 am 
		greetingMsg = "Good night"
	}

	return greetingMsg

}


func main(){

  fmt.Println("Current time is : ",time.Now().Format(time.Kitchen))
  //kitchen is built-in time layout defined in the time package 
  //example of kitchen format is 3:04PM

  msg := greeting()
  fmt.Println(msg)

}
