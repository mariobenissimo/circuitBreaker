package internal

type BreakerState interface {
	nextState() BreakerState
	isClosed() bool
}
