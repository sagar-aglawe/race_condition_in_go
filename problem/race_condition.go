package problem

import (
	"fmt"
	"sync"
)

var x = 0

func increase(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func ProblemDemo() {

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increase(&wg)
	}
	wg.Wait()

	fmt.Println(fmt.Sprintf("Value of X is with race condition is := %v", x))
}
