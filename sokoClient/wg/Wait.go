package wg

import "time"

func WaitAsyncOf(check func() bool, callback func()) {
	for {
		if check() {
			callback()
			break
		}
		time.Sleep(time.Second / 10)
	}
}
