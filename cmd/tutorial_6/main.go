package main

import "fmt"

func main() {
	// pointers
	//var p *int32 = new(int32)
	//var i int32
	//fmt.Printf("The value of p pointer to is: %v\n", *p)
	//fmt.Printf("The value of i is: %v\n", i)
	//*p = 10
	//fmt.Printf("The value of p pointer to is: %v\n", *p)
	//p = &i // now p points to i so the out put is 0
	//fmt.Printf("The value of p pointer to is: %v\n", *p)
	//
	//var slice = []int32{1, 2, 3, 4, 5} //pointer so the out put is the same for slice and sliceCopy
	//var sliceCopy = slice
	//sliceCopy[2] = 4
	//fmt.Println(slice)
	//fmt.Println(sliceCopy)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\n the memory location of thing1 is: %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Println(result)
	fmt.Printf("The value of thing1 is: %v\n", thing1)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\n The memory location of thing2 array is: %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2 // now since we used the pointer in the square as the parameter things 1 and thing 2 are the same memory location
	// so the out put will be the same
}
