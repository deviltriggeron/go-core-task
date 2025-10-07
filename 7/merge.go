package mergechan

import "sync"

func Produce(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i <= n; i++ {
			ch <- i
		}
	}()
	return ch
}

func MergeChan(chs ...<-chan int) chan int {
	res := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(chs))
	for _, ch := range chs {
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				res <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}
