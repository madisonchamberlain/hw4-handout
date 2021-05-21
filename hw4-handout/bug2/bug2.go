package bug2

import (
	"sync"
)

func bug2(n int, foo func(int) int, out chan int) {
	var wg sync.WaitGroup
	// make a channel of size n
	ch := make(chan int, n)
	for i := 0; i < n; i++ {
		// put i into the channel
		ch <- i
		wg.Add(1)
		go func() {
			// output the channel contents 
			out <- foo(<-ch)
			wg.Done()
		}()
	}
	wg.Wait()
	close(out)
}
