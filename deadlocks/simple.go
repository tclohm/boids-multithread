package main

import (
	"sync"
	"fmt"
	"time"
)

var (
	lock1 = sync.Mutex{}
	lock2 = sync.Mutex{}
)

func bluePhilosopher() {
	for {
		fmt.Println("Blue is acquiring lock1")
		lock1.Lock()
		fmt.Println("Blue is acquiring lock2")
		lock2.Lock()
		fmt.Println("Blue acquired both locks")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Blue has released locks")
	}
}

func redPhilosopher() {
	for {
		fmt.Println("Red is acquiring lock2")
		lock2.Lock()
		fmt.Println("Red is acquiring lock1")
		lock1.Lock()
		fmt.Println("Red acquired both locks")
		lock1.Unlock()
		lock2.Unlock()
		fmt.Println("Red has released locks")
	}
}

func main() {
	go redPhilosopher()
	go bluePhilosopher()
	time.Sleep(20 * time.Second)
	fmt.Println("done")
}