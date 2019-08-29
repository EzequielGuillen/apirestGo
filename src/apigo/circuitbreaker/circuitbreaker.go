package circuitbreaker

import "time"

type CircuitBreaker struct {
	State      string
	CantErrors int
	TimeOut    time.Duration
}

func (circuit *CircuitBreaker) SetState(state string)  {
	circuit.State=state
}
