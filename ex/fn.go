package ex

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func say(msg string) {
	fmt.Println(msg)
}

func count(max int, ch chan int) {

	for i := 0; i < max; i++ {
		ch <- i
	}
}

func countWithClose(max int, ch chan int) {

	for i := 0; i < max; i++ {
		ch <- i
	}
	close(ch)
}

func chReceiver(ch, quit chan int) {
	//
	// get 3 then tell the sender to quit
	//
	for i := 0; i < 3; i++ {
		fmt.Printf("got: %v\n", <-ch)
	}
	quit <- 0

}

func countWithQuit(ch, quit chan int) {
	i := 0

	for {
		select {
		case ch <- i:
			fmt.Printf("sent: %v\n", i)
			i++
		case <-quit:
			fmt.Println("Quit - got signle on the quit channel")
			close(ch)
			return
		}
	}

}

func sendFast(ch chan int) {

	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("sent: %v\n", i)
	}
	close(ch)
}
func sendSlow(ch chan int) {

	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("sent: %v\n", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func recSlow(ch chan int) {

	for i := range ch {
		fmt.Printf("got: %v\n", i)
		time.Sleep(100 * time.Millisecond)
	}

}

func recFast(ch chan int) {

	for i := range ch {
		fmt.Printf("got: %v\n", i)
	}

}

func sayWG(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("say: %s\n", msg)
}

func wasteTime(wg *sync.WaitGroup) {
	defer wg.Done()
	i := 123
	s := fmt.Sprintf("Just wasting time on %d", i)
	s = s + strconv.Itoa(i)
}
