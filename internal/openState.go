package internal

import (
	"fmt"
	"time"
)

type OpenState struct {
	openTime      int64
	deltaTimeOpen int64
	str           string
}

func NewOpenState(delta int64, name string) *OpenState {
	currentTime := (time.Now().Unix()) * 100
	fmt.Println("got to " + name)
	return &OpenState{
		openTime:      currentTime,
		deltaTimeOpen: delta,
		str:           name,
	}
}

func nextStateOpen(op *Open) BreakerState {
	elapsed := ((time.Now().Unix()) * 100) - op.openTime
	fmt.Println("time in " + op.str + " " + fmt.Sprint(elapsed) + "ms, ")
	if elapsed < op.deltaTimeOpen {
		fmt.Println("stay " + op.str)
		return op // resto nello stesso stao
	}
	return NewHalfOpen()
}
func nextStateHalf(op *HalfOpen) BreakerState {
	elapsed := ((time.Now().Unix()) * 100) - op.openTime
	fmt.Println("time in " + op.str + " " + fmt.Sprint(elapsed) + "ms, ")
	if elapsed < op.deltaTimeOpen {
		fmt.Println("stay " + op.str)
		return op // resto nello stesso stao
	}
	return NewClosed()
}
func (op *OpenState) nextState() BreakerState { // metodi fake per ora
	return op
}
func (op *OpenState) isClosed() bool { // metodi fake per ora
	return false
}
