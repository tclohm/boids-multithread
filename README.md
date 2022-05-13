### Process, Threads, and Green Threads

ready queue, context switching from operating system (context switch overhead)
interrupts occur possible reading from disc

#### Process
Own Application with own memory space --- isolation
--- Someone draws a house on a piece of paper
--- Someone draws a the land and sky on a piece of paper
--- Someone draws a tree on a piece of paper
--- Combine papers together to create scene


#### Thread
Application that shares memory space -- faster than process because not allocating memory
--- Same Paper but everyone has to communicate on what they are drawing and where they are drawing
--- Make sure threads don't step on each other
--- When the operating systems swaps out a process or thread and chooses another one, from the run
	queue, to be executed -- this is called Context Switching
--- When 2 or more threads call lock(), on a non-locked mutex, at the same time...exactly one thread
	acquires the lock

#### Green Thread
More efficient
--- user level thread
--- operating system doesnt know anything from it

##### Memory Sharing
--- Memory sharing is a manner in which threads can communicate with each other
	This is possible because threads created from the same process, share the 
	same memory space

##### Golang uses a hybrid system Thread and Green Thread 

##### Locks
Solution: Each thread is exclusively locked for the entire loop, so it prints all the numbers first
before it releases the lock and another thread acquires it
1,2,3,1,2,3

If we used a sync.RWMutex{}, with a RLock() and RUnlock(), 
readers lock allows multiple threads to execute the same read locked block of code it's impossible to know the exact order that the threads will be interleaved
go`
    var lock = sync.Mutex{}
    func oneTwoThreeA() {
    	lock.Lock()
    	for i := 1; i <= 3; i++ {
    		fmt.Println(i)
    		time.Sleep(1 * time.Millisecond)
    	}
    	lock.Unlock()
    }
    func StartThreadsA() {
    	for i := 1; i <= 2; i++ {
    		go oneTwoThreeA()
    	}
    	time.Sleep(1 * time.Second)
    }
`

Calling the lock twice in a row, will make the call block forever. 
The mutexexs in GO are not re-entrant though. So! The thread already on lock
is not affected and would return immediately --- Thanks GO!

#### Messages

In go, channels are a type of thread/process communication (IPC) that uses message passing between threads

go`
    func runConsumer(channel chan string) {
       msg := <-channel
       fmt.Println("Consumer, received", msg)
       channel <- "Bye"
    }
    func RunProducer() {
       channel := make(chan string)
       go runConsumer(channel)
       fmt.Println("Producer Sending Hello")
       channel <- "Hello"
       fmt.Println("Producer, received", <-channel)
    }
`

output: Producer Sending Hello
		Consumer, received Hello
		Producer, received Bye


#### Buffer Channel

go`func BuffSender() {
   channel := make(chan string, 3)
   fmt.Println("Sending ONE")
   channel <- "ONE"
   fmt.Println("Sending TWO")
   channel <- "TWO"
   fmt.Println("Sending THREE")
   channel <- "THREE"
   fmt.Println("Done")
    }`
Solution: Sending ONE, Sending TWO, Sending THREE, Done

#### Wait Groups
--- Provides another way to do thread sync
go`
func waitGroupQueuesA() {
	wg := sync.WaitGroup{}
	wg.Wait()
	fmt.Println("Done")
}
`
solution: it will output "Done" immediately because we are not calling Add(number) on the WG

go`
func count() {
   wg := sync.WaitGroup{}
   x := 0
   wg.Add(5)
   for i := 0; i < 5; i ++ {
      go increment(&x, &wg)
   }
   wg.Wait()
   fmt.Printf("%d\n",x)
}
func increment(x *int, wg *sync.WaitGroup) {
   for i := 0; i < 100; i ++ {
      *x += 1
   }
   wg.Done()
}
`
Solution: It's impossible to tell as the function has a race condition
The main thread creates 5 threads and waits for them to complete
Each thread adds 100 to the counter
However threads might overwrite each other, since access to x is not
exclusive

##### Conditional
thread calls wait() on a conditional variable it will unlock the mutex 
associated with that condition variable

a condition variable provides a signal() function
only one thread is unblocked. Upon unblocking the thread will try 
to reacquire the lock associated with the conditional variable

go`
    func runChildThread() {
       mlock.Lock()
       fmt.Println("RunChildThread, lock acquired")
       cond.Signal()
       fmt.Println("RunChildThread, Waiting")
       cond.Wait()
       fmt.Println("RunChildThread, Running")
    }
    func RunMainThread() {
       mlock.Lock()
       fmt.Println("RunMainThread, lock acquired")
       go runChildThread()
       fmt.Println("RunMainThread, Waiting")
       cond.Wait()
       fmt.Println("RunMainThread, Running")
       cond.Signal()
       time.Sleep(10 * time.Second)
    }
`
executes: RunMainThread, lock acquired
RunMainThread, Waiting
RunChildThread, lock acquired
RunChildThread, Waiting
RunMainThread, Running