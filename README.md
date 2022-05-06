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

#### Green Thread
More efficient
--- user level thread
--- operating system doesnt know anything from it

##### Golang uses a hybrid system Thread and Green Thread 