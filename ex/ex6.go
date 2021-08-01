package ex

import "time"

func Ex6Main() {

	ch := make(chan int,7) // the larger the buffer the faster the 
	                       // send routine finishes

	go recSlow(ch)
	sendFast(ch)

	time.Sleep(2 * time.Second)
}
