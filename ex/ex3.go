package ex

import "fmt"

//
// Ex3Main - Buildes on Ex2 but will uses a range to receive from ch
//
func Ex3Main() {
	ch := make(chan int)

	go countWithClose(10, ch)

	for i := range ch {
		fmt.Printf("Received from ch: %v\n", i)
	}
	
}


