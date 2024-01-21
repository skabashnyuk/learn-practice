// calculate maximum value in slice
package main
import (
	"fmt"
)
func main (){
	nums := []int{ 16,8,42,4,23,15}
	max:=0
	for _, i := range nums {
		if i>max {
			max=i
		}
	}
	fmt.Println(max)
}
