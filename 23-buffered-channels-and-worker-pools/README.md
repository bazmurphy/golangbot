# Buffered Channels and Worker Pools

## What are buffered channels?

All the channels we discussed in the previous tutorial were basically unbuffered. As we discussed in the channels tutorial in detail, sends and receives to an unbuffered channel are blocking.

It is possible to create a channel with a buffer. Sends to a buffered channel are blocked only when the buffer is full. Similarly receives from a buffered channel are blocked only when the buffer is empty.

Buffered channels can be created by passing an additional capacity parameter to the `make` function which specifies the size of the buffer.

```go
ch := make(chan type, capacity)
```

`capacity` in the above syntax should be greater than 0 for a channel to have a buffer.
The capacity for an unbuffered channel is 0 by default and hence we omitted the capacity parameter while creating channels in the previous tutorial.

Let’s write some code and create a buffered channel.

```go
package main

import (
	"fmt"
)

func main() {
	// make a buffered channel (the 2 is the capacity, we can write 2 strings to the channel)
	ch := make(chan string, 2)
	// send a string to the channel
	ch <- "naveen"
	// send another string to the channel
	ch <- "paul"
	// read from the channel and print
	fmt.Println(<-ch)
	// read from the channel and print
	fmt.Println(<-ch)
}
```

In the program above, in line no. 9 we create a buffered channel with a capacity of 2. Since the channel has a capacity of 2, it is possible to write 2 strings into the channel without being blocked. We write 2 strings to the channel in line no. 10 and 11 and the channel does not block. We read the 2 strings written in line nos. 12 and 13 respectively. This program prints,

```
naveen
paul
```

## Another Example

Let’s look at one more example of the buffered channel in which the values to the channel are written in a concurrent Goroutine and read from the main Goroutine. This example will help us better understand when writes to a buffered channel block.

```go
package main

import (
	"fmt"
	"time"
)

// takes a channel as an argument
func write(ch chan int) {
	for i := 0; i < 5; i++ {
		// sends i to the channel
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	// closes the channel
	close(ch)
}

func main() {
	// make a new channel of type int with capacity 2 (buffered channel)
	ch := make(chan int, 2)
	// run a Goroutine
	go write(ch)
	time.Sleep(2 * time.Second)
	// receive the values from the channel
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		// sleep for 2 seconds, this blocks the main Goroutine and any further reading from the channel
		time.Sleep(2 * time.Second)
	}
}

```

In the program above, a buffered channel `ch` of capacity `2` is created in line no. 16 of the `main` Goroutine and passed to the `write` Goroutine in line no. 17. Then the main Goroutine sleeps for 2 seconds. During this time, the `write` Goroutine is running concurrently. The `write` Goroutine has a `for` loop which writes numbers from 0 to 4 to the `ch` channel. The capacity of this buffered channel is `2` and hence the write `Goroutine` will be able to write values `0` and `1` to the `ch` channel immediately and then it blocks until at least one value is read from `ch` channel. So this program will print the following 2 lines immediately.

```
successfully wrote 0 to ch
successfully wrote 1 to ch
```

After printing the above two lines, the writes to the `ch` channel in the `write` Goroutine are blocked until someone reads from the `ch` channel. Since the main Goroutine sleeps for 2 seconds before starting to read from the channel, the program will not print anything for the next 2 seconds. The `main` Goroutine wakes up after 2 seconds and starts reading from the `ch` channel using a `for range` loop in line no. 19, prints the read value and then sleeps for 2 seconds again and this cycle continues until the `ch` is closed. So the program will print the following lines after 2 seconds

```
read value 0 from ch
successfully wrote 2 to ch
```

This will continue until all values are written to the channel and it is closed in the write Goroutine. The final output would be,

```
successfully wrote 0 to ch
successfully wrote 1 to ch
read value 0 from ch
successfully wrote 2 to ch
read value 1 from ch
successfully wrote 3 to ch
read value 2 from ch
successfully wrote 4 to ch
read value 3 from ch
read value 4 from ch
```

## Deadlock

```go
package main

import (
	"fmt"
)

func main() {
	// we make a buffered channel of capacity 2
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	// so when the control reaches the third write
	// the write is blocked since the channel has exceeded its capacity
	ch <- "steve"
	// some Goroutine needs to read from the channel to unblock it so the write can proceed
	// so there is a deadlock and the program panics
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
// fatal error: all goroutines are asleep - deadlock!
// goroutine 1 [chan send]:
```

In the program above, we write 3 strings to a buffered channel of capacity 2. When the control reaches the third write in line no. 11, the write is blocked since the channel has exceeded its capacity. Now some Goroutine must read from the channel in order for the write to proceed, but in this case, there is no concurrent routine reading from this channel. Hence there will be a deadlock and the program will panic at run time with the following message,

## Closing buffered channels

We already discussed about closing channels in the previous tutorial. In addition to what we have learned in the previous tutorial, there is one more subtlety to be considered when closing buffered channels.

It’s possible to read data from a already closed buffered channel. The channel will return the data that is already written to the channel and once all the data has been read, it will return the zero value of the channel.

Let’s write a program to understand this.

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 5)
	ch <- 5
	ch <- 6
	close(ch)
	n, open := <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d, open: %t\n", n, open)

	// Received: 5, open: true
	// Received: 6, open: true
	// Received: 0, open: false
}
```

In the program above, we created a buffered channel of capacity `5` in line no. 8. We then write `5` and `6` to the channel. The channel is closed after that in line no. 11. Even though the channel is closed, we can read the values already written to the channel. This is done in line nos. 12 and 14. The value of `n` will be `5` and open will be `true` in line no. 12. The value of `n` will be `6` and `open` will be `true` again in line no 14. We have now finished reading 5 and 6 from the channel and there is no more data to be read. Now when the channel is read again in line no. 16, the value of `n` will be `0` which is the zero value of `int` and `open` will be `false` indicating that the channel has closed.

```
Received: 5, open: true
Received: 6, open: true
Received: 0, open: false
```

The same program can be written using for range loop too.

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 5)
	ch <- 5
	ch <- 6
	close(ch)
	// The for range loop will read all the values written to the channel
	// and will quit once there are no more values to read since the channel is already closed.
	for n := range ch {
		fmt.Println("Received:", n)
	}
}

// Received: 5
// Received: 6
```

The `for range` loop in line no. 12 of the program above will read all the values written to the channel and will quit once there are no more values to read since the channel is already closed.

This program will print,

```
Received: 5
Received: 6
```

## Length vs Capacity

The capacity of a buffered channel is the number of values that the channel can hold. This is the value we specify when creating the buffered channel using the `make` function.

The length of the buffered channel is the number of elements currently queued in it.

A program will make things clear

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 3)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read value", <-ch)
	fmt.Println("new length is", len(ch))
}

// capacity is 3
// length is 2
// read value naveen
// new length is 1
```

In the program above, the channel is created with a capacity of 3, that is, it can hold 3 strings. We then write 2 strings to the channel in line nos. 9 and 10 respectively. Now the channel has 2 strings queued in it and hence its length is 2. In line no. 13, we read a string from the channel. Now the channel has only one string queued in it and hence its length becomes 1. This program will print,

```
capacity is 3
length is 2
read value naveen
new length is 1
```

## WaitGroup

The next section in this tutorial is about Worker Pools. To understand worker pools, we need to first know about WaitGroup as it will be used in the implementation of Worker pool.

A WaitGroup is used to wait for a collection of Goroutines to finish executing. The control is blocked until all Goroutines finish executing. Let’s say we have 3 concurrently executing Goroutines spawned from the main Goroutine. The main Goroutines needs to wait for the 3 other Goroutines to finish before terminating. This can be accomplished using WaitGroup.

Let’s stop the theory and write some code right away

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

// started Goroutine  2
// started Goroutine  1
// started Goroutine  0
// Goroutine 0 ended
// Goroutine 2 ended
// Goroutine 1 ended
// All go routines finished executing
```

`WaitGroup` is a struct type and we are creating a zero value variable of type `WaitGroup` in line no.18. The way `WaitGroup` works is by using a counter. When we call `Add` on the `WaitGroup` and pass it an `int`, the `WaitGroup`’s counter is incremented by the value passed to `Add`. The way to decrement the counter is by calling `Done()` method on the WaitGroup. The `Wait()` method blocks the Goroutine in which it’s called until the counter becomes zero.

In the above program, we call `wg.Add(1)` in line no. 20 inside the `for` loop which iterates 3 times. So the counter now becomes 3. The `for` loop also spawns 3 `process` Goroutines and then `wg.Wait()` called in line no. 23 makes the main Goroutine to wait until the counter becomes zero. The counter is decremented by the call to `wg.Done` in the process Goroutine in line no. 13. Once all the 3 spawned Goroutines finish their execution, that is once `wg.Done()` has been called three times, the counter will become zero, and the main Goroutine will be unblocked.

It is important to pass the pointer of `wg` in line no. 21. If the pointer is not passed, then each Goroutine will have its own copy of the `WaitGroup` and main will not be notified when they finish executing.

This program outputs.

```
started Goroutine 2
started Goroutine 0
started Goroutine 1
Goroutine 0 ended
Goroutine 2 ended
Goroutine 1 ended
All go routines finished executing
```

Your output might be different from mine since the order of execution of Goroutines can vary

## Worker Pool Implementation

One of the important uses of buffered channel is the implementation of [worker pool](https://en.wikipedia.org/wiki/Thread_pool) (thread pool).

In general, a worker pool is a collection of threads that are waiting for tasks to be assigned to them. Once they finish the task assigned, they make themselves available again for the next task.

We will implement a worker pool using buffered channels. Our worker pool will carry out the task of finding the sum of a digits of the input number. For example if 234 is passed, the output would be 9 (2 + 3 + 4). The input to the worker pool will be a list of pseudo-random integers.

The following are the core functionalities of our worker pool

- Creation of a pool of Goroutines which listen on an input buffered channel waiting for jobs to be assigned
- Addition of jobs to the input buffered channel
- Writing results to an output buffered channel after job completion
- Read and print results from the output buffered channel

We will write this program step by step to make it easier to understand.

The first step will be the creation of the structs representing the job and the result.

```go
type Job struct {
	id       int
	randomno int
}

type Result struct {
	job         Job
	sumofdigits int
}
```

Each `Job` struct has a `id` and a `randomno` for which the sum of the individual digits has to be computed.

The `Result` struct has a `job` field which is the job for which it holds the result (sum of individual digits) in the `sumofdigits` field.

The next step is to create the buffered channels for receiving the jobs and writing the output.

```go
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)
```

Worker Goroutines listen for new tasks on the `jobs` buffered channel. Once a task is complete, the result is written to the `results` buffered channel.

The `digits` function below does the actual job of finding the sum of the individual digits of an integer and returning it. We will add a sleep of 2 seconds to this function just to simulate the fact that it takes some time for this function to calculate the result.

```go
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}
```

Next, we will write a function that creates a worker Goroutine.

```go
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}
```

The above function creates a worker which reads from the `jobs` channel, creates a `Result` struct using the current `job` and the return value of the `digits` function and then writes the result to the `results` buffered channel. This function takes a WaitGroup `wg` as a parameter on which it will call the `Done()` method when all `jobs` have been completed.

The `createWorkerPool` function will create a pool of worker Goroutines.

```go
func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}
```

The function above takes the number of workers to be created as a parameter. It calls `wg.Add(1)` before creating the Goroutine to increment the WaitGroup counter. Then it creates the worker Goroutines by passing the pointer of the WaitGroup `wg` to the `worker` function. After creating the needed worker Goroutines, it waits for all the Goroutines to finish their execution by calling `wg.Wait()`. After all Goroutines finish executing, it closes the `results` channel since all Goroutines have finished their execution and no one else will further be writing to the `results` channel.

Now that we have the worker pool ready, let’s go ahead and write the function which will allocate jobs to the workers.

```go
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}
```

The `allocate` function above takes the number of jobs to be created as input parameter, generates pseudo random numbers with a maximum value of `998`, creates `Job` struct using the random number and the for loop counter `i` as the id and then writes them to the `jobs` channel. It closes the `jobs` channel after writing all jobs.

The next step would be to create the function that reads the `results` channel and prints the output.

```go
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}
```

The `result` function reads the `results` channel and prints the job id, input random no, and the sum of digits of the random no. The result function also takes a `done` channel as a parameter to which it writes to once it has printed all the results.

We have everything set now. Let’s go ahead and finish the last step of calling all these functions from the `main()` function.

```go
func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
```

We first store the execution start time of the program in line no.2 of the main function and in the last line (line no. 12) we calculate the time difference between the endTime and startTime and display the total time it took for the program to run. This is needed because we will do some benchmarks by changing the number of Goroutines.

The `noOfJobs` is set to 100 and then `allocate` is called to add jobs to the `jobs` channel.

Then `done` channel is created and passed to the `result` Goroutine so that it can start printing the output and notify once everything has been printed.

Finally a pool of `10` worker Goroutines are created by the call to `createWorkerPool` function and then main waits on the `done` channel for all the results to be printed.

Here is the full program for your reference. I have imported the necessary packages too.

```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}

type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
```

Please run this program in your local machine for more accuracy in the total time taken calculation.

This program will print,

```
Job id 1, input random no 636, sum of digits 15
Job id 0, input random no 878, sum of digits 23
Job id 9, input random no 150, sum of digits 6
...
total time taken  20.01081009 seconds
```

A total of 100 lines will be printed corresponding to the 100 jobs and then finally the total time taken for the program to run will be printed in the last line. Your output will differ from mine as the Goroutines can run in any order and the total time will also vary based on the hardware. In my case, it takes approximately 20 seconds for the program to complete.

Now let’s increase the `noOfWorkers` in the `main` function to `20`. We have doubled the number the workers. Since the worker Goroutines have increased(doubled to be precise), the total time taken for the program to complete should reduce(by half to be precise). In my case it became, 10.004364685 seconds and the program printed,

```
total time taken  10.004364685 seconds
```

Now we can understand that as the number of worker Goroutines increases, the total time taken to complete the jobs decreases. I leave it as an exercise for you to play with the `noOfJobs` and `noOfWorkers` in the `main` function to different values and analyze the results.
