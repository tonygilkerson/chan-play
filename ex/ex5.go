package ex

import (
	"fmt"
	"time"
)

//
// Ex5Main - A variation on Ex4, this time we are using a buffer ch
//
func Ex5Main() {
	ch := make(chan int, 5) // If you add a buffer you will send more than you receive
	quit := make(chan int)

	go chReceiver(ch, quit)
	go countWithQuit(ch, quit)

	// the sender was able to send more than what the receiver wanted because of the ch buffer
	// lets see if we can slurp up the left over message on the ch
	time.Sleep(2 * time.Second)
	fmt.Printf("\nTry to slurp...\n")
	for i := range ch {
		fmt.Printf("slurp got %v\n", i)
	}

}
