package ex

import (
	"fmt"
	"time"
)

func Ex8Main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(3 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Print(" ")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Print("_")
			time.Sleep(30 * time.Millisecond)
		}
	}
}
