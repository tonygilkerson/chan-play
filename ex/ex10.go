package ex

import (
	"fmt"
	"runtime"
	"sync"
)

func Ex10Main() {

	ncpu := 4
	fmt.Printf("\nSetting ncpu to: %d, it was %d\n", ncpu, runtime.GOMAXPROCS(ncpu))

	wg := new(sync.WaitGroup)

	for i := 0; i < 100000000; i++ {
		wg.Add(1)
		go wasteTime(wg)
	}

	fmt.Printf("Start wait\n")
	wg.Wait()
	fmt.Println("Done waiting")

}
