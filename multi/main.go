package main

import (
	"sync"
	"fmt"
	"time"
)

// Since the readers lock allows multiple threads to execute 
// the same read locked block of code it's impossible to know 
// the exact order that the threads will be interleaved.

var rwlock = sync.RWMutex{}
 
func oneTwoThreeB() {
	rwlock.RLock()
	for i := 1; i <= 300; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Millisecond)
	}
	rwlock.RUnlock()
}
 
func StartThreadsB() {
	for i := 1; i <= 2; i++ {
		go oneTwoThreeB()
	}
	time.Sleep(1 * time.Second)
}

func main() {
	StartThreadsB()
}