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

var (
	matrixA = [matrixSize][matrixSize]int{}
	matrixB = [matrixSize][matrixSize]int{}
	result = [matrixSize][matrixSize]int{}
	lck = sync.RWMutex{}
	cond = sync.NewCond(lck.RLocker())
	wg := sync.WaitGroup()
)

func generateRandom(matrix *[matrixSize][matrixSize]int) {
	lck.RLock()
	for {
		wg.Done()
		cond.Wait()
		for row := 0 ; row < matrixSize ; row++ {
			for col := 0 ; col < matrixSize ; col++ {
				matrix[row][col] += rand.Intn(10) - 5
			}
		}
	}
	
}

func workoutmatrix(row int) {
	wg.Add(matrixSize)
	for row := 0 ; row < matrixSize ; row++ {
		// thread
		go workoutmatrix(row)
	}

	for i := 0 ; i < 100 ; i++ {
		wg.Wait()
		lck.Lock()
		generateRandom(&matrixA)
		generateRandom(&matrixB)
		wg.Add(matrixSize)
	}
	lck.Unlock()
	cond.Broadcast()
}

func runit() {
	for row := 0 ; row < matrixSize ; row++ {
		workoutmatrix(row)
	}
}