package utils

import (
	"../circuitbreaker"
	"fmt"
	"net/http"
	"time"
)

var (
	Client = http.Client{
		Timeout: 3*time.Second,
	}
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
		fmt.Println("Si")

		response, err :=Client.Get(UrlCountryPing)
		if err != nil || response.StatusCode == 500 {
			continue
		}

		response, err =Client.Get(UrlUserPing)
		if err != nil || response.StatusCode == 500 {
			continue
		}
		response, err = Client.Get(UrlSitePing)
		if err != nil || response.StatusCode ==  500{
			continue
		}
		CircuitBreaker.State = "CLOSE"
		return
	}
}
