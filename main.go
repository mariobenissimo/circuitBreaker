package main

import (
	"fmt"
	"time"

	"github.com/mariobenissimo/circuitBreaker/internal"
)

func main() {
	brk := internal.NewCircuitBreaker()
	var result int
	for i := 0; i < 10; i++ {
		result = brk.GetResultBoundTime()
		fmt.Println("Iteration: " + fmt.Sprint(i+1) + " result: " + fmt.Sprint(result))
		time.Sleep(1 * time.Second)
	}
}
