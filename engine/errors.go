package engine

import "errors"

var (
	ErrMaxIteration  = errors.New("exceeded max iteration")
	ErrFatalFeedback = errors.New("fatal feedback encountered")
)
