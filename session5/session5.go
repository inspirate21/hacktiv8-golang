package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

var breakpoint = 13

func main() {
	c := make(chan string)
	quit := make(chan string)
	// go random(4, c)
	go random(4, c, quit)

	// msg1 := <-c
	fmt.Println(<-c)

	// msg2 := <-c
	// fmt.Println("player 2 ", msg2)
}

func random(totalPlayer int, c chan string) {
	var no int
	var result int
	for i := 1; i <= totalPlayer; i++ {
		min := 1
		max := 100
		no = rand.Intn(max-min) + min
		result = no % breakpoint
		if result == 0 {
			fmt.Println("player ke", i, "dapat angka", no, "dengan breakpoint", breakpoint, "sehingga permainan berhenti")

			c <- strconv.Itoa(result)
			return
		} else {
			fmt.Println("player ke", i, "dapat angka", no, "dengan breakpoint", breakpoint, "sehingga permainan dilanjutkan")
		}

		if i == totalPlayer {
			i = 0
		}
	}
}

// func randomChannel(totalPlayer int, c, quit chan string) {
// 	var no int
// 	var result int
// 	for i := 1; i <= totalPlayer; i++ {
// 		min := 1
// 		max := 100
// 		no = rand.Intn(max-min) + min
// 		result = no % breakpoint
// 		select {
// 		case result == 0:
// 			// x, y = y, x+y
// 			fmt.Println("player ke", i, "dapat angka", no, "dengan breakpoint", breakpoint, "sehingga permainan berhenti")
// 		case <-quit:
// 			fmt.Println("quit")
// 			return
// 		}

// 		// if result == 0 {
// 		// 	fmt.Println("player ke", i, "dapat angka", no, "dengan breakpoint", breakpoint, "sehingga permainan berhenti")

// 		// 	c <- strconv.Itoa(result)
// 		// 	return
// 		// } else {
// 		// 	fmt.Println("player ke", i, "dapat angka", no, "dengan breakpoint", breakpoint, "sehingga permainan dilanjutkan")
// 		// }

// 		if i == totalPlayer {
// 			i = 0
// 		}
// 	}
// }
