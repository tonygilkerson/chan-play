package ex

import (
	"time"
)

//
// Ex4Main - Tell the sender to quit
//
func Ex4Main() {
	// ch := make(chan int,12) // If you add a buffer you will send more than you receive
	ch := make(chan int)
	quit := make(chan int)

	go chReceiver(ch, quit)
	go countWithQuit(ch, quit)

	time.Sleep(3 * time.Second)
}
