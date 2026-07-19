package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString string = "résumé"
	println(myString)
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("\n the lenget of my string is %v \n", len(myString))

	var myString1 = []rune("résumé")
	var indexed1 = myString1[1]
	fmt.Printf("%v, %T\n", indexed1, indexed1)
	for a, b := range myString1 {
		fmt.Println(a, b)
	}

	fmt.Printf("\n the lenget of my string is %v \n", len(myString1))

	//	╰─ go run string.go
	//		résumé
	//	114, uint8
	//	0 114
	//	1 233
	//	3 115
	//	4 117
	//	5 109
	//	6 233
	//
	//	the lenget of my string is 8
	//	233, int32
	//	0 114
	//	1 233
	//	2 115
	//	3 117
	//	4 109
	//	5 233
	//
	//	the lenget of my string is 6

	var strSlice = []string{"s", "u", "b"}
	var carStr = ""
	for i := range strSlice {
		carStr += strSlice[i]
	}
	fmt.Println(carStr)

	var strSlice1 = []string{"s", "u", "b"}
	var strBuilder strings.Builder
	for i := range strSlice1 {
		strBuilder.WriteString(strSlice1[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("%v", catStr)

}
