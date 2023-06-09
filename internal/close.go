package internal

import "fmt"

type Closed struct{}

func NewClosed() *Closed {
	fmt.Println("go to closed")
	return &Closed{}
}

func (c *Closed) nextState() BreakerState {
	return c
}

func (c *Closed) isClosed() bool {
	return true
}
