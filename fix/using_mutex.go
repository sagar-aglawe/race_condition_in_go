package fix

import (
	"fmt"
	"sync"
)

var y = 0

func increaseUsingMutex(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	y = y + 1
	wg.Done()
	m.Unlock()
}

func FixUsingMutex() {

	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increaseUsingMutex(&wg, &m)
	}
	wg.Wait()

	fmt.Println(fmt.Sprintf("Value of Y post using mutex is := %v", y))
}
