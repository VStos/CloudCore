package tools

import "time"

func Rest(duration time.Duration) {

	done := make(chan bool)
	go func() {
		time.Sleep(duration)
		done <- true
	}()
}
