package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Control_flow() {

	// if condition {
	// 	// code to be executed if condition is true
	//   }

	rand.Seed(time.Now().UnixNano()) // needed before Go 1.20

	if n := rand.Int(); n%2 == 0 {
		fmt.Println(n, "is an even number.")
	} else {
		fmt.Println(n, "is an odd number.")
	}

	n := rand.Int() % 2 // this n is not the above n.

	if n % 2 == 0 {
		fmt.Println("An even number.")
	}

	if ; n % 2 != 0 {
		fmt.Println("An odd number.")
	}
}