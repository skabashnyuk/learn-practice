package main
import (
	"fmt"
)
/*
Print the number between 1 to 20, however
 - if the number is devisible by 3, print "fizz" instead
 - if the number is devisible by 5, print "buzz" instead
 - if the number is devisible by 3 and 5, print "fizz buzz" instead
 - otherwise print the number
*/


func main (){
	// fmt.Println(1%3)	
	for i:=1; i<=20; i++ {
//	fmt.Printf("-----%v", i)
	fmt.Println("---")
	fmt.Println(i)
//	fmt.Println(i%3)
//	fmt.Println(i%5)
	fmt.Println("----")
		switch{
			case i%15==0:
			  fmt.Println("fizz buzz")
			case i%3==0:
			  fmt.Println("fizz")
			case i%5==0:
			  fmt.Println("buzz") 
			default:
			  fmt.Println(i)    
		}
	}
}
