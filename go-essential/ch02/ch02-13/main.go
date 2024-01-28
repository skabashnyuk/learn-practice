/*
 count how many times each word appears in a text.
 This is a very basic step in many text
 processing algorithms.

 To spit text to words, use the "Firelds" function
 from the "stings" package.

 Also use "ToLower" form the same package
 to conver words to lowercase
 
 */
 package main

 import (
 	"fmt"
 	"strings"
 )

 func main(){
    text:=`
 count how many times each word appears in a text.
 This is a very basic step in many text
 processing algorithms.

 To spit text to words, use the "Firelds" function
 from the "stings" package.

 Also use "ToLower" form the same package
 to conver words to lowercase
 `
	flds := strings.Fields(text)
	countMap := map[string] int {}
	for _,s := range flds {
		//fmt.Printf("s %v type %T", s,s)
		//fmt.Println("")
		fld := strings.ToLower(s)
		count:=countMap[fld]
		countMap[fld]=count+1
	}
 	fmt.Printf("flds %v with type %T", countMap, countMap)
 }
