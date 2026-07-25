package main

import "fmt"

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner // alos can use just owner and use gasEngine.name in the main
}

type owner struct {
	name string
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// struct method
func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("can make it")
	} else {
		fmt.Println("cannot make it")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15, owner{"alex"}}
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)

	var myEngine2 = struct { //anonymous struct . we cant reuse it. have to create new one
		mpg       uint32
		gallons   uint32
		ownerInfo owner
	}{10, 15, owner{"indika"}}
	fmt.Println(myEngine2.mpg, myEngine2.gallons, myEngine2.ownerInfo.name)

	//calling the method of struct
	fmt.Println("Total miles left: ", myEngine.milesLeft())

	canMakeIt(myEngine, 20)

	var myEngine4 electricEngine = electricEngine{25, 15}
	canMakeIt(myEngine4, 50)
}
