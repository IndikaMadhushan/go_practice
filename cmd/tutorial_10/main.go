package main

import "fmt"

func main() {
	//generics
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))
}

func sumSlice[T int | float64 | float32](slice []T) T { // can use this for accept any type [T any]
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
