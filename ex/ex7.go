package ex

import (
	"time"
)

func Ex7Main() {

	// A larger buffer has no affect because of the slow sender. 
	// i.e. the receiver receives as fast as it can send
	ch := make(chan int, 5)

	go recFast(ch)
	sendSlow(ch)

	time.Sleep(2 * time.Second)
}
