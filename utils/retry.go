package utils

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	RetryDuration = time.Second
	RetryMaxTimes = 5
)

func Retry[T any](do func() (T, bool, error)) (T, error) {
	return RetryWithTimes(RetryMaxTimes, do)
}

func RetryWithTimes[T any](times int, do func() (T, bool, error)) (T, error) {
	for i := 0; i < times; i++ {
		ret, retry, err := do()
		if !retry {
			return ret, err
		}

		log.Printf("failed to do: %s, retry %d", err, i)
		time.Sleep(RetryDuration)
	}
	log.Printf("exceed max retry times, drop")
	var ret T
	return ret, fmt.Errorf("exceed max retry times")
}

func RetryRequest(req *http.Request) (*http.Response, error) {

	return Retry(func() (*http.Response, bool, error) {
		var httpClient = http.Client{}
		var res *http.Response
		res, err := httpClient.Do(req)
		if err != nil {
			return nil, true, err
		}
		return res, false, err
	})
}
