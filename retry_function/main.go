package main

import (
	"errors"
	"fmt"
	"log"
)

func doRetry(attempts int, f func() error) error {
	log.Printf("call function failed. attempts=%d\n", attempts)

	if attempts <= 0 {
		return errors.New("error: try the best")
	}

	if err := f(); err != nil {
		attempts--
		return doRetry(attempts, f)
	}

	log.Println("call function success")
	return nil
}

func alwaysFailed() error {
	log.Println("do alwaysFailed function")
	return errors.New("call alwaysFailed: always failed")
}

// for succInThirdCall, record call count
var callCount int

func succInThirdCall() error {
	callCount++
	log.Println("do call succInThirdCall, will success in third call")

	if callCount < 3 {
		return fmt.Errorf("call function count = %d failed", callCount)
	}

	return nil
}

func main() {
	if err := doRetry(3, alwaysFailed); err != nil {
		log.Printf("doRetry error: %v", err)
	}

	if err := doRetry(3, succInThirdCall); err != nil {
		log.Printf("doRetry error: %v", err)
	}
}
