package internal

import (
	"fmt"
	"reflect"
	"time"
)

type CircuitBreaker struct {
	t     Task
	state BreakerState
}

func NewCircuitBreaker() *CircuitBreaker {
	return &CircuitBreaker{
		t:     *newTask(),
		state: NewClosed(),
	}
}
func (c *CircuitBreaker) GetResultBoundTime() int {
	c.state = c.state.nextState()
	fmt.Println(reflect.TypeOf(c.state))
	if !c.state.isClosed() {
		return -1 // se già è aperto possiamo restituire -1
	}
	fmt.Println("CircuitBreaker executing task") // proviamo a fare una richiesta
	resultCh := make(chan int8)
	errorCh := make(chan error)
	timeout := time.After(300 * time.Millisecond)
	go func() {
		value, err := c.t.NextStep() // next step del task
		if err != nil {
			fmt.Println("An exception occurred in Task")
			errorCh <- err
		} else {
			resultCh <- value
		}
	}()

	select {
	case result := <-resultCh:
		fmt.Println("Async request completed successfully. Result:", result)
		return int(result)
	case err := <-errorCh:
		fmt.Println("Async request failed with error:", err)
		c.state = NewOpen()
		fmt.Println(reflect.TypeOf(c.state))
		return -1
	case <-timeout:
		c.state = NewOpen()
		return -1
	}
}
