package main

import (
	"time"
	"fmt"
	"sync"
)

var (
	money = 100
	lock = sync.Mutex{}
	moneyDepo = sync.NewCond(&lock)
)

func stingy() {
	for i := 1 ; i <= 1000 ; i++ {
		lock.Lock()
		money += 10
		fmt.Println("Stingy balance", money)
		// If conditional is false, we can start spending again
		// signal will unlock the conditional lock
		moneyDepo.Signal()
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Stingy Done")
}

func spendy() {
	for i := 1 ; i <= 1000 ; i++ {
		lock.Lock()
		for money - 20 < 0 {
			moneyDepo.Wait()
		}
		money -= 20
		fmt.Println("Spendy balance", money)
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Spendy Done")
}

func main() {
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println(money)
}