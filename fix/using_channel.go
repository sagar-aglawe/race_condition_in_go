package fix

import (
	"fmt"
	"sync"
)

var z = 0

func increaseUsingChannel(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	z = z + 1
	wg.Done()
	<-ch
}

func FixUsingChannel() {

	var wg sync.WaitGroup
	ch := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increaseUsingChannel(&wg, ch)
	}
	wg.Wait()

	fmt.Println(fmt.Sprintf("Value of Z post using channel is := %v", z))
}
