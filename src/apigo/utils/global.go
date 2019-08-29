package utils

import (
	"../circuitbreaker"
	"net/http"
	"time"
)

var (
	CircuitBreaker = circuitbreaker.CircuitBreaker{
		State:      "CLOSE",
		CantErrors: 3,
		TimeOut:    500 * time.Millisecond,
	}
)

func TimeOut() {
	for {
		CircuitBreaker.State = "OPEN"
		time.Sleep(CircuitBreaker.TimeOut)
		CircuitBreaker.State = "HALFOPEN"
		response, err := http.Get(UrlCountryPing)
		if err != nil || response.StatusCode == 500 {
			continue
		}
		response, err = http.Get(UrlUserPing)
		if err != nil || response.StatusCode == 500 {
			continue
		}
		response, err = http.Get(UrlSitePing)
		if err != nil || response.StatusCode ==  500{
			continue
		}
		CircuitBreaker.State = "CLOSE"
		return
	}
}
