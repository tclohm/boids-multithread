package main

import (
	"sync"
	"runtime"
	"sync/atomic"
)

type SpinLock int32

func (s *SpinLock) Lock() {
	// casting s as an integer
	for !atomic.CompareAndSwapInt32((*int)(s), 0, 1) {
		runtime.Gosched()
	}
}

func (s *SpinLock) Unlock() {
	atomic.StoreInt32((*int32)(s), 0)
}

func NewSpinLock() sync.Locker {
	var lock SpinLock
	return &lock
}