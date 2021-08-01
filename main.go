package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/tonygilkerson/chan-play/ex"
)

func main() {
	exampleNo, _ := strconv.Atoi(os.Args[1])

	fmt.Printf("EXAMPLE [%v] - Start...\n", exampleNo)

	start := time.Now()
	switch exampleNo {
	case 1:
		ex.Ex1Main()
	case 2:
		ex.Ex2Main()
	case 3:
		ex.Ex3Main()
	case 4:
		ex.Ex4Main()
	case 5:
		ex.Ex5Main()
	case 6:
		ex.Ex6Main()
	case 7:
		ex.Ex7Main()
	case 8:
		ex.Ex8Main()
	case 9:
		ex.Ex9Main()
	case 10:
		ex.Ex10Main()		
	default:
		fmt.Println("Default case, get out.")
	}

	duration := time.Since(start)
	fmt.Printf("EXAMPLE [%v] - Done, took: %v\n", exampleNo, duration)

}
