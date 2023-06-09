package internal

import (
	"fmt"
	"math/rand"
)

type HalfOpen struct {
	OpenState
}

func NewHalfOpen() *HalfOpen {
	return &HalfOpen{
		OpenState: *NewOpenState(100, "half open"),
	}
}
func (ho *HalfOpen) nextState() BreakerState {
	return nextStateHalf(ho)
}
func (ho *HalfOpen) isClosed() bool {
	fmt.Println("Closed di halfopen")
	randomNumber := rand.Intn(2)
	if randomNumber == 0 {
		return true
	} else {
		return false
	}
}
