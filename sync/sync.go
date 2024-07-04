package main

import "sync"

type Counter struct {
	mu  sync.Mutex
	Val int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Val++
}

func (c *Counter) Value() int {
	return c.Val
}

func NewCounter() *Counter {
	return &Counter{}
}
