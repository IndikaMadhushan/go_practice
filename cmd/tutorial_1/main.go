package main

import (
	"errors"
	"fmt"
)

func main() {
	//fmt.Println("fmt")
	//var intNum int = 3276788
	//fmt.Println(intNum)

	//var floatNum float64 = 12345678.9
	//fmt.Println(floatNum)

	//var floatNum32 float32 = 10.1
	//var intNum32 int32 = 2
	//var result float32 = floatNum32 + float32(intNum32)
	//fmt.Println(result)

	//var myString string = "Hello " + " " + "World"
	//fmt.Println(myString)
	//
	//fmt.Println(utf8.RuneCountInString("a"))
	//
	//var myRune rune = 'a'
	//fmt.Println(myRune)
	//
	//var myBool bool = true
	//fmt.Println(myBool)
	//
	//var intNum4 rune
	//fmt.Println(intNum4)

	//myVar := "text"
	//fmt.Println(myVar)
	//
	//var var1, var2 = 1, 3
	//fmt.Println(var1, var2)

	//const myConst string = "hello"
	//fmt.Println(myConst)

	printMe("new tesxt")
	var firstValue, second, err = intDivision(10, 6)
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else if second == 0 {
	//	fmt.Printf("the result is %v", firstValue)
	//
	//} else {
	//	fmt.Printf("the result is %v and remainder is %v \n", firstValue, second)
	//}
	//fmt.Println(firstValue, second, err)

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case second == 0:
		fmt.Println("the result is ", firstValue)
	default:
		fmt.Println("the result is ", firstValue, " and remainder is ", second)

	}

	//var myInt int = 10
	//var myFloat float64 = 10.1
	//fmt.Printf("first value is %v and second value is %v", myInt, myFloat)
}

func printMe(printValue string) {
	fmt.Println("I am a ")
}

func intDivision(numerator int, denomerator int) (int, int, error) {
	var err error
	if denomerator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}

	var result int = numerator / denomerator
	var remainder int = numerator % denomerator
	return result, remainder, err // err = nil
}
