package internal

import "fmt"

type Open struct {
	OpenState
}

func NewOpen() *Open {
	return &Open{
		OpenState: *NewOpenState(200, "open"),
	}
}
func (o *Open) nextState() BreakerState {
	//fmt.Println("QUI1", reflect.TypeOf(o))
	return nextStateOpen(o)
}
func (o *Open) isClosed() bool {
	fmt.Println("Closed di open")
	return false
}
