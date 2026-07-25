package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
	var chickenChannel = make(chan string)
	var website = []string{"chicken.com", "tofu.com", "wallmark.com"}
	for i := range website {
		go checkChickenPrice(website[i], chickenChannel)
	}
}

func checkChickenPrice(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice < MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string) {
	fmt.Printf("\n Found a deal on chicken at %s\n", <-chickenChannel)
}
