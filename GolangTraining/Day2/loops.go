package main

import "fmt"

func main(){
	i:=5
	for{
		fmt.Println("Hello ",i)
		i += 5
		if i == 20 {
			break
		}
	}

	for j := 0 ; j<10 ; j++ {
		println("Index increasing : ",j)
	} 

	
	for j := 10 ; j>5 ; j-- {
		println("Index decreasing : ",j)
	} 

	sum := 0
	for ; sum<5 ; sum=sum+2{    //Need not write sum := 0 again 
		println("Sum value : ",sum)
	}

}



// Only FOR loop in Golang
//to break outer for loop you can use FOR LABEL
//myFor:
//for i:=0;i<10;i++{
//	for j:=0;j<10;j++{
//		 if j==1{
//			 break myFor
//       }
//    }
//}