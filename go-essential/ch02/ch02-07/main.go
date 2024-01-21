/* 
An 'even ended number' is a number whose first and last digit are the same.

You mission, should you choose to accept it, is to count how many 'even ended numbers' are
there that are a muliplication of two 4 diggit numbers.
*/

package main
import (
	"fmt"
)

func main (){
   t:=0
	for i:=1000; i<=9999; i++ {
		for j:=i; j<=9999; j++ {
			s:=fmt.Sprintf("%d",i*j)
			if s[0]==s[len(s)-1]{
				t++
				fmt.Println(s)
			}		
		}
	}
	fmt.Printf("Total: %v \n",t)
}
