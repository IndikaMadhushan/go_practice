package main

import "fmt"

func main() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"adam": 23, "john": 45}
	fmt.Println(myMap2)

	fmt.Println("hel")

	fmt.Println(myMap2["adam"])
	fmt.Println(myMap2["mala"])

	var age, ok = myMap2["adam"]
	fmt.Println(age, ok)

	if ok {
		fmt.Println("age is ", age)
	} else {
		fmt.Println("invalid name")
	}

	for name, age := range myMap2 {
		fmt.Println("name is ", name, " and age is ", age)
	}

	var intArr = [5]int{1, 2, 3, 4, 5}

	for i, v := range intArr {
		fmt.Println(i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
