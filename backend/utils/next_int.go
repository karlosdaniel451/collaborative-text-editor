package utils

import "sync"

func NextIntGenerator(start int) func() int {
	value := start
	lock := sync.Mutex{}
	return func() int {
		lock.Lock()
		defer func ()  {
			value++	
			lock.Unlock()
		}()
		return value
	}
}