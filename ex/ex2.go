package ex

import "fmt"

//
// Ex2Main - use a channel to get the results from a go routine
//
func Ex2Main() {
	ch := make(chan int)

	go count(3, ch)

	fmt.Printf("First  Received from ch: %v\n", <-ch)
	fmt.Printf("Second Received from ch: %v\n", <-ch)
	fmt.Printf("Third  Received from ch: %v\n", <-ch)

	fmt.Printf("This will cause deadlock!,  Received from ch: %v\n", <-ch)
	
}


