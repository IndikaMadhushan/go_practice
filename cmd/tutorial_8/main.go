package main

import (
	"fmt"
	"time"
)

// unbuffered channel var c = make(chan int)
// buffered channel var c = make(chan int, 10)
func main() {
	// channels hold data
	// thread safe
	// listen for data
	//var c = make(chan int)
	var c = make(chan int, 5) // now the process function can add 5 numbers to the buffer and finish without waiting for main
	//c <- 1   this is a deadlock
	//var i = <-c
	//fmt.Println(i)

	go process(c)
	for i := range c { // if the channel is not closed the loop will go for 6 th value and get a deadlock
		fmt.Println(i)
		time.Sleep(time.Second * 1) // adding a delay in the main
	}
}
func process(c chan int) { // channel need to close otherwise can encounter deadlock
	defer close(c) // close the channel when the function is done
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("exiting process")

}
