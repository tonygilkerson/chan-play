package ex

import (
	"fmt"
	"strconv"
	"time"
)

//
// Ex1Main - Simple launching of go routines
//
func Ex1Main() {

	for i := 0; i < 10; i++ {
		go say("this is " + strconv.Itoa(i))
	}

	// If we dont wait we will quit the program before we
	// see the output of the go routines.
	// There is a better way!
	time.Sleep(3 * time.Second)
	fmt.Println("Done.")
}
