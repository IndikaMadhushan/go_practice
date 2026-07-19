package main

import "fmt"

func main() {
	var intArr [3]int32 = [3]int32{1, 2, 3} // fixed length
	//intArr:= [...]int32{1, 2, 3}

	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	var intSlice []int32 = []int32{1, 2, 3}
	fmt.Println(intSlice)
	fmt.Println(len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 4)
	fmt.Println(intSlice)
	fmt.Println(len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 5)
	intSlice = append(intSlice, 6)
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Println(len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9, 10}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 = make([]int32, 0, 5)
	fmt.Println(intSlice3)

}
