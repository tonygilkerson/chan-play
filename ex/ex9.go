package ex

import (
	"fmt"
	"strconv"
	"sync"
)

func Ex9Main() {

	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go sayWG(strconv.Itoa(i), wg)
	}

	fmt.Println("start waiting...")
	wg.Wait()
	fmt.Println("done waiting")
}
