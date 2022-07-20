/* 
objects of a superclass should be replaceable with objects of its subclasses without breaking the application

so we have Mobile embedded inside samsung and nokia
inside SendText(m) Mobile m is passed
SendText(n1) nokia is passed
so interchangeably both objects can be passed

*/

package main

import "fmt"

type texter interface {
	SendText()
}
type Mobile struct {
	Price int
}

func (m Mobile) SendText() {
	fmt.Println("Sending text from mobile")
}

type Samsung struct{
	Mobile
}

type Nokia struct{
	Mobile
}

func SendText(t texter){
	t.SendText()
}

func main(){
	s1 := Samsung{}
	SendText(s1)
	n1 := Nokia{}
	SendText(n1)
	m := Mobile{}
	m.SendText()
	SendText(m)

}