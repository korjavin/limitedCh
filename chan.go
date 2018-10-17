package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("vim-go")
}

type limitedChan struct {
	ch chan int
}

func newLimitedChan(len int) *limitedChan {
	l := &limitedChan{
		ch: make(chan int, len),
	}
	go l.worker()
	return l
}
func (c *limitedChan) worker() {
	for range c.ch {
		time.Sleep(time.Minute)
	}
}
func (c *limitedChan) WriteWithLen(msg int) {
	if len(c.ch) >= cap(c.ch) {
		return
	}
	c.ch <- msg
}
func (c *limitedChan) WriteWithDefault(msg int) {
	select {
	case c.ch <- msg:
	default:
	}
}
