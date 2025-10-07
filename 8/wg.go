package mywg

import "time"

type CustomWg struct {
	ch chan struct{}
}

func NewCustomWg() *CustomWg {
	return &CustomWg{
		ch: make(chan struct{}, 10000),
	}
}

func (wg *CustomWg) Add(n int) {
	for i := 0; i < n; i++ {
		wg.ch <- struct{}{}
	}
}

func (wg *CustomWg) Done() {
	<-wg.ch
}

func (wg *CustomWg) Wait() {
	for len(wg.ch) > 0 {
		time.Sleep(1 * time.Millisecond)
	}
}
